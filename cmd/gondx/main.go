package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/dashboard"
	"github.com/bonedaddy/go-indexed/db"
	"github.com/bonedaddy/go-indexed/discord"
	"github.com/bonedaddy/go-indexed/utils"
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
		&cli.BoolFlag{
			Name:  "startup.sleep",
			Usage: "whether or not to sleep on startup, useful for giving containers time to initialize",
			Value: false,
		},
		&cli.DurationFlag{
			Name:  "startup.sleep_time",
			Usage: "time.Duration type specifying sleep duration",
			Value: time.Second * 5,
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.Bool("startup.sleep") {
			time.Sleep(c.Duration("startup.sleep_time"))
		}
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
			Name:  "prometheus",
			Usage: "serves the prometheus metric endpoint",
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
				defer database.Close()
				if err := database.AutoMigrate(); err != nil {
					return err
				}
				// serve prometheus metrics
				go dashboard.ServePrometheusMetrics(ctx, c.String("listen.address"))
				// update metrics information
				go dashboard.UpdateMetrics(ctx, database, bc)
				sc := make(chan os.Signal, 1)
				signal.Notify(sc, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt, os.Kill)
				<-sc
				cancel()
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "listen.address",
					Usage: "the address we will expose the metrics listener on",
					Value: "0.0.0.0:9123",
				},
			},
		},
		&cli.Command{
			Name:  "discord",
			Usage: "discord bot management",
			Subcommands: cli.Commands{
				&cli.Command{
					Name:    "database",
					Aliases: []string{"db"},
					Usage:   "database management commands",
					Subcommands: cli.Commands{
						&cli.Command{
							Name:        "chain-updater",
							Usage:       "starts the database chain state updater",
							Description: "chain-updater is responsible for persisting all information needed into a database such as price updates",
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
								defer database.Close()
								if err := database.AutoMigrate(); err != nil {
									return err
								}
								wg := &sync.WaitGroup{}
								wg.Add(1)
								// launch database price updater loop
								go func() {
									defer wg.Done()
									dbPriceUpdateLoop(ctx, bc, database)
								}()

								sc := make(chan os.Signal, 1)
								signal.Notify(sc, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt, os.Kill)
								<-sc
								cancel()
								wg.Wait()
								return nil
							},
						},
					},
				},
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
						defer database.Close()
						if err := database.AutoMigrate(); err != nil {

							return err
						}
						wg := &sync.WaitGroup{}
						if c.Bool("update.database") {
							wg.Add(1)
							// launch database price updater loop
							go func() {
								defer wg.Done()
								dbPriceUpdateLoop(ctx, bc, database)
							}()
						}

						client, err := discord.NewClient(ctx, cfg, bc, database)
						if err != nil {
							return err
						}
						sc := make(chan os.Signal, 1)
						signal.Notify(sc, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt, os.Kill)
						<-sc
						cancel()
						wg.Wait()
						return client.Close()
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "discord.token",
							Usage:   "the discord api token",
							EnvVars: []string{"DISCORD_TOKEN"},
						},
						&cli.BoolFlag{
							Name:  "update.database",
							Usage: "if true launch the db price update routine. if false make sure chain-updater command is running",
							Value: true,
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
	ticker := time.NewTicker(time.Second * 60) // update every 1m
	defer ticker.Stop()
	// update TVL every 15 minutes
	tvlTicker := time.NewTicker(time.Minute * 15)
	defer tvlTicker.Stop()

	// update the tvl price before looping to insert some data
	go func() {
		cc10, err := bc.CC10()
		if err != nil {
			log.Println("failed to get cc10 for tvl price update: ", err)
			return
		}
		defi5, err := bc.DEFI5()
		if err != nil {
			log.Println("failed to get defi5 for tvl price udpate: ", err)
			return
		}
		cc10TVL, err := bc.GetTotalValueLocked(cc10)
		if err != nil {
			log.Println("failed to get total value locked for cc10: ", err)
			return
		}
		defi5TVL, err := bc.GetTotalValueLocked(defi5)
		if err != nil {
			log.Println("failed to get total value locked for defi5: ", err)
			return
		}
		if err := db.RecordValueLocked("defi5", defi5TVL); err != nil {
			log.Println("failed to update tvl for defi5: ", err)
			return
		}
		if err := db.RecordValueLocked("cc10", cc10TVL); err != nil {
			log.Println("failed to update tvl for cc10: ", err)
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tvlTicker.C:
			cc10, err := bc.CC10()
			if err != nil {
				log.Println("failed to get cc10 for tvl price update: ", err)
			} else {
				cc10TVL, err := bc.GetTotalValueLocked(cc10)
				if err != nil {
					log.Println("failed to get total value locked for cc10: ", err)
				} else {
					if err := db.RecordValueLocked("cc10", cc10TVL); err != nil {
						log.Println("failed to update tvl for cc10: ", err)
					}
				}
			}

			defi5, err := bc.DEFI5()
			if err != nil {
				log.Println("failed to get defi5 for tvl price udpate: ", err)
			} else {
				defi5TVL, err := bc.GetTotalValueLocked(defi5)
				if err != nil {
					log.Println("failed to get total value locked for defi5: ", err)
				} else {
					if err := db.RecordValueLocked("defi5", defi5TVL); err != nil {
						log.Println("failed to update tvl for defi5: ", err)
					}
				}
			}

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
			// update defi5 total supply
			ip, err := bc.DEFI5()
			if err != nil {
				log.Println("failed to get defi5 contract: ", err)
			} else {
				supply, err := ip.TotalSupply(nil)
				if err != nil {
					log.Println("failed to get total supply: ", err)
				} else {
					supplyF, _ := utils.ToDecimal(supply, 18).Float64()
					if err := db.RecordTotalSupply("defi5", supplyF); err != nil {
						log.Println("failed to update defi5 total supply")
					}
				}
			}
			// update cc10 total supply
			ip, err = bc.CC10()
			if err != nil {
				log.Println("failed to get cc10 contract: ", err)
			} else {
				supply, err := ip.TotalSupply(nil)
				if err != nil {
					log.Println("failed to get total supply: ", err)
				} else {
					supplyF, _ := utils.ToDecimal(supply, 18).Float64()
					if err := db.RecordTotalSupply("cc10", supplyF); err != nil {
						log.Println("failed to update cc10 total supply")
					}
				}
			}
			// update ndx total supply
			erc, err := bc.NDX()
			if err != nil {
				log.Println("failed to get ndx contract: ", err)
			} else {
				supply, err := erc.TotalSupply(nil)
				if err != nil {
					log.Println("failed to get total supply: ", err)
				} else {
					supplyF, _ := utils.ToDecimal(supply, 18).Float64()
					if err := db.RecordTotalSupply("ndx", supplyF); err != nil {
						log.Println("failed to update ndx total supply")
					}
				}
			}
		}

	}
}
