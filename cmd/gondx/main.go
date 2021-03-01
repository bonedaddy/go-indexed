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
	"github.com/bonedaddy/go-indexed/config"
	"github.com/bonedaddy/go-indexed/dashboard"
	"github.com/bonedaddy/go-indexed/db"
	"github.com/bonedaddy/go-indexed/discord"
	"github.com/bonedaddy/go-indexed/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
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
		return nil
	}
	app.Commands = cli.Commands{
		&cli.Command{
			Name:  "prometheus",
			Usage: "serves the prometheus metric endpoint",
			Action: func(c *cli.Context) error {
				ctx, cancel := context.WithCancel(c.Context)
				defer cancel()
				cfg, err := config.LoadConfig(c.String("config"))
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
								cfg, err := config.LoadConfig(c.String("config"))
								if err != nil {
									return err
								}
								logger, err := config.LoggerFromConfig(cfg)
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
								// do an initial seed
								processUpdates(bc, database, logger, cfg.Indices)
								// launch database price updater loop
								go func() {
									defer wg.Done()
									dbPriceUpdateLoop(ctx, bc, database, logger, cfg)
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
						return config.NewConfig(c.String("config"))
					},
				},
				&cli.Command{
					Name:  "ndx-bot",
					Usage: "starts NDXBot",
					Action: func(c *cli.Context) error {
						ctx, cancel := context.WithCancel(c.Context)
						defer cancel()
						cfg, err := config.LoadConfig(c.String("config"))
						if err != nil {
							return err
						}
						logger, err := config.LoggerFromConfig(cfg)
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
							// do an initial seed
							processUpdates(bc, database, logger, cfg.Indices)
							wg.Add(1)
							// launch database price updater loop
							go func() {
								defer wg.Done()
								dbPriceUpdateLoop(ctx, bc, database, logger, cfg)
							}()
						}

						client, err := discord.NewClient(ctx, cfg, bc, database, logger)
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
							Value: false,
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

func dbPriceUpdateLoop(ctx context.Context, bc *bclient.Client, db *db.Database, logger *zap.Logger, cfg *config.Config) {
	ticker := time.NewTicker(time.Second * 240) // update every 4m
	defer ticker.Stop()
	// update TVL every 35 minutes
	tvlTicker := time.NewTicker(time.Minute * 35)
	defer tvlTicker.Stop()

	// do an initial update
	processUpdates(bc, db, logger, cfg.Indices)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			processUpdates(bc, db, logger, cfg.Indices)
		}
	}

}

// handleIndexUpdates is responsible for handling updates of indice
// tvl, total supply, and price
func handleIndexUpdates(db *db.Database, bc *bclient.Client, logger *zap.Logger, indices []string) {
	for _, indice := range indices {
		curr := strings.ToLower(indice)
		ip, err := bc.GetIndexPool(curr)
		if err != nil {
			break
		}
		supply, price, tvl, err := getValues(bc, ip, logger, curr)
		if err := db.RecordValueLocked(curr, tvl); err != nil {
			logger.Error("failed to record value locked", zap.Error(err), zap.String("indice", indice))
		}
		if err := db.RecordPrice(curr, price); err != nil {
			logger.Error("failed to record price", zap.Error(err), zap.String("indice", indice))
		}
		if err := db.RecordTotalSupply(curr, supply); err != nil {
			logger.Error("failed to record total supply", zap.Error(err), zap.String("indice", indice))
		}
	}
}

func getValues(bc *bclient.Client, ip bclient.IndexPool, logger *zap.Logger, name string) (totalSupply, price, tvl float64, err error) {
	totalSupply, err = getTotalSupply(ip)
	switch name {
	case "defi5":
		price, err = bc.Defi5DaiPrice()
	case "cc10":
		price, err = bc.Cc10DaiPrice()
	case "orcl5":
		price, err = bc.Orcl5DaiPrice()
	case "degen10":
		price, err = bc.Degen10DaiPrice()
	}
	tvl, err = bc.GetTotalValueLocked(ip, logger)
	return
}

func getTotalSupply(erc bclient.ERCO20I) (float64, error) {
	totalSupplyBig, err := erc.TotalSupply(nil)
	if err != nil {
		return 0, err
	}
	totalSupplyF, _ := utils.ToDecimal(totalSupplyBig, 18).Float64()
	return totalSupplyF, nil
}

func processUpdates(bc *bclient.Client, db *db.Database, logger *zap.Logger, indices []string) {
	// do an initial update
	handleIndexUpdates(db, bc, logger, indices)
	// update eth price seperately as it requires special handling
	priceBig, err := bc.EthDaiPrice()
	if err != nil {
		logger.Error("failed to get eth dai price", zap.Error(err))
	} else {
		if err := db.RecordPrice("eth", float64(priceBig.Int64())); err != nil {
			logger.Error("failed to update eth dai price", zap.Error(err))
		}
	}
	// update ndx price seperately as it requires special handling
	price, err := bc.NdxDaiPrice()
	if err != nil {
		logger.Error("failed to get ndx dai price", zap.Error(err))
	} else {
		if err := db.RecordPrice("ndx", price); err != nil {
			logger.Error("failed to udpate ndx dai price", zap.Error(err))
		}
	}

	// update ndx total supply seperately as it requires special handling
	erc, err := bc.NDX()
	if err != nil {
		logger.Error("failed to get ndx contract", zap.Error(err))
	} else {
		supply, err := erc.TotalSupply(nil)
		if err != nil {
			logger.Error("failed to get total supply", zap.Error(err))
		} else {
			supplyF, _ := utils.ToDecimal(supply, 18).Float64()
			if err := db.RecordTotalSupply("ndx", supplyF); err != nil {
				logger.Error("failed to udpate ndx total supply", zap.Error(err))
			}
		}
	}

}
