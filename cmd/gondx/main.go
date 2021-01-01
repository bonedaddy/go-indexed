package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/discord"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var (
	bc *bclient.Client
)

func main() {
	app := cli.NewApp()
	app.Name = "gondx"
	app.Usage = "a CLI application for Indexed Finance"
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
						cfg, err := discord.LoadConfig(c.String("config"))
						if err != nil {
							return err
						}
						if bc != nil {
							if cfg.InfuraAPIKey != "" {
								bc, err = bclient.NewInfuraClient(cfg.InfuraAPIKey, cfg.InfuraWSEnabled)
							} else {
								bc, err = bclient.NewClient(cfg.ETHRPCEndpoint)
							}
							if err != nil {
								return err
							}
							defer bc.Close()
						}
						client, err := discord.NewClient(cfg, bc)
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
