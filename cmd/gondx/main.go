package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/db"
	"github.com/bonedaddy/go-indexed/discord"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var (
	bc      *bclient.Client
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = "gondx"
	app.Usage = "a CLI application for Indexed Finance"
	app.Version = Version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "infura.api_key",
			Usage:   "api key for use with infura",
			EnvVars: []string{"INFURA_API_KEY"},
		},
		&cli.StringFlag{
			Name:  "eth.rpc",
			Usage: "specifies the ethereum RPC endpoint",
			Value: bclient.InfuraHTTPURL,
		},
		&cli.StringFlag{
			Name:  "eth.address",
			Usage: "address to lookup in queries",
			Value: "0x5a361A1dfd52538A158e352d21B5b622360a7C13",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"cfg"},
			Usage:   "path to discord bot configuration file",
			Value:   "config.yml",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.String("config") != "" {
			return nil
		}
		var err error
		if c.String("infura.api_key") != "" {
			var websockets bool
			if strings.Contains(c.String("infura.api_key"), "wss") {
				websockets = true
			}
			bc, err = bclient.NewInfuraClient(c.String("infura.api_key"), websockets)
		} else {
			bc, err = bclient.NewClient(c.String("eth.rpc"))
		}
		return err
	}
	app.Commands = cli.Commands{
		&cli.Command{
			Name:  "discord",
			Usage: "discord bot management",
			Subcommands: cli.Commands{
				&cli.Command{
					Name:  "gen-config",
					Usage: "generate ndx bot config file",
					Action: func(c *cli.Context) error {
						return discord.NewConfig(c.String("config"))
					},
				},
				&cli.Command{
					Name:  "ndx-bot",
					Usage: "starts NDXBot",
					Action: func(c *cli.Context) error {
						ctx, cancel := context.WithCancel(c.Context)
						defer cancel()
						cfg, err := discord.LoadConfig(c.String("config"))
						if err != nil {
							return err
						}
						if cfg.InfuraAPIKey != "" {
							bc, err = bclient.NewInfuraClient(cfg.InfuraAPIKey, cfg.InfuraWSEnabled)
						} else {
							bc, err = bclient.NewClient(cfg.ETHRPCEndpoint)
						}
						if err != nil {
							return err
						}
						defer bc.Close()
						database, err := db.New(&db.Opts{
							Type:           cfg.Database.Type,
							Host:           cfg.Database.Host,
							Port:           cfg.Database.Port,
							User:           cfg.Database.User,
							Password:       cfg.Database.Pass,
							DBName:         cfg.Database.DBName,
							SSLModeDisable: cfg.Database.SSLModeDisable,
						})
						if err != nil {
							return err
						}
						if err := database.AutoMigrate(); err != nil {
							database.Close()
							return err
						}
						defer database.Close()
						// launch database price updater loop
						go dbPriceUpdateLoop(ctx, bc, database)
						client, err := discord.NewClient(ctx, cfg, bc, database)
						if err != nil {
							return err
						}
						sc := make(chan os.Signal, 1)
						signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
						<-sc
						return client.Close()
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "discord.token",
							Usage:   "the discord api token",
							EnvVars: []string{"DISCORD_TOKEN"},
						},
					},
				},
				&cli.Command{
					Name:  "price-bot",
					Usage: "starts a price monitoring bot",
					Action: func(c *cli.Context) error {
						return errors.New("not yet implemented")
					},
				},
			},
		},
		&cli.Command{
			Name:  "pool",
			Usage: "index pool commands",
			Subcommands: cli.Commands{
				&cli.Command{
					Name:  "balance-of",
					Usage: "returns decimal amount of pool tokens held by address",
					Action: func(c *cli.Context) error {
						var (
							ip  bclient.IndexPool
							err error
						)
						switch c.String("pool.name") {
						case "defi5":
							ip, err = bc.DEFI5()
						case "cc10":
							ip, err = bc.CC10()
						default:
							return errors.New("unsupported index pool")
						}
						bal, err := bclient.BalanceOfDecimal(ip, common.HexToAddress(c.String("eth.address")))
						if err != nil {
							return err
						}
						fmt.Println(bal)
						return nil
					},
				},
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "pool.name",
					Usage: "the name of the pool to lookup",
					Value: "defi5",
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func dbPriceUpdateLoop(ctx context.Context, bc *bclient.Client, db *db.Database) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// update defi5 price
			price, err := bc.Defi5DaiPrice()
			if err != nil {
				log.Println("failed to get defi5 dai price: ", price)
			} else {
				if err := db.RecordPrice("defi5", price); err != nil {
					log.Println("failed to update defi5 dai price: ", err)
				}
			}
			// update cc10 price
			price, err = bc.Cc10DaiPrice()
			if err != nil {
				log.Println("failed to get cc10 dai price: ", err)
			} else {
				if err := db.RecordPrice("cc10", price); err != nil {
					log.Println("failed to update cc10 dai price: ", err)
				}
			}
			// update eth price
			priceBig, err := bc.EthDaiPrice()
			if err != nil {
				log.Println("failed to get eth dai price: ", err)
			} else {
				if err := db.RecordPrice("eth", float64(priceBig.Int64())); err != nil {
					log.Println("failed to update eth dai price: ", err)
				}
			}
			// update ndx price
			price, err = bc.NdxDaiPrice()
			if err != nil {
				log.Println("failed to get ndx dai price: ", err)
			} else {
				if err := db.RecordPrice("ndx", price); err != nil {
					log.Println("failed to update ndx dai price: ", err)
				}
			}
		}

	}
}
