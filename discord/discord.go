package discord

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/config"
	"github.com/bonedaddy/go-indexed/db"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"golang.org/x/text/message"
)

var (
	rateLimitMsg = "You are being rate limited. Users are allowed 1 blockchain query per command every 60 seconds"
	printer      *message.Printer
)

// Client wraps bclient and discordgo to provide a discord bot for indexed finance
type Client struct {
	s  *discordgo.Session
	bc *bclient.Client
	r  *dgc.Router
	db *db.Database

	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup

	logger *zap.Logger

	cfg *config.Config
}

func init() {
	printer = message.NewPrinter(message.MatchLanguage("en"))
}

// NewClient provides a wrapper around discordgo
func NewClient(ctx context.Context, cfg *config.Config, bc *bclient.Client, db *db.Database, logger *zap.Logger) (*Client, error) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(ctx)
	// launch price watcher bots
	if err := launchWatchers(ctx, wg, cfg, bc, db, logger); err != nil {
		cancel()
		return nil, err
	}
	// register the watchers
	dg, err := discordgo.New("Bot " + cfg.MainDiscordToken)
	if err != nil {
		cancel()
		return nil, err
	}

	if err := dg.Open(); err != nil {
		cancel()
		return nil, err
	}

	// updates the main bot's nickname to reflect ndx price
	ndxPriceWatchRoutine(ctx, dg, wg, db, logger)

	// declare the base router
	router := dgc.Create(&dgc.Router{
		Prefixes:         []string{"!ndx"},
		IgnorePrefixCase: true,  // allow any case of the prefix
		BotsAllowed:      false, // do not allow bots
		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("ping")
		},
	})

	client := &Client{s: dg, bc: bc, r: router, ctx: ctx, cancel: cancel, wg: wg, db: db, cfg: cfg}

	// register our custom help command
	registerHelpCommand(dg, nil, router)

	router.RegisterCmd(&dgc.Command{
		Name:        "info",
		Description: "provides information about NDXBot",
		IgnoreCase:  true,
		Handler: func(ctx *dgc.Ctx) {
			base := BaseEmbed()
			base.Title = "Indexed Finance Bot Information"
			base.Description = "NDXBot is an unofficial, open-source discord bot that faciltates making read-only queries to the Indexed Finance protocol"
			base.Fields = []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:  "Source Code",
					Value: "https://github.com/bonedaddy/go-indexed/tree/main/discord",
				},
				&discordgo.MessageEmbedField{
					Name:  "Bot Documentation",
					Value: "https://github.com/bonedaddy/go-indexed/blob/main/docs/DISCORD_BOT.md",
				},
				&discordgo.MessageEmbedField{
					Name:  "Indexed Finance Documentation",
					Value: "https://docs.indexed.finance/",
				},
				&discordgo.MessageEmbedField{
					Name:  "Support",
					Value: "star the [github repo](https://github.com/bonedaddy/go-indexed), or [tip some ethereum based cryptos](https://etherscan.io/address/0x5a361a1dfd52538a158e352d21b5b622360a7c13). all tipped funds will be reinvested into Indexed Finance",
				},
			}
			ctx.RespondEmbed(base)
		},
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "pool",
		Description: "command group for interacting with Indexed pools",
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "current-tokens",
				Description: "returns the current tokens basketed into a pool",
				Usage:       " pool current-tokens <pool-name>",
				Example:     " pool current-tokens defi5",
				IgnoreCase:  true,
				Handler:     client.poolTokensHandler,
				// We want the user to be able to execute this command once in 60 seconds and the cleanup interval shpuld be one second
				RateLimiter: dgc.NewRateLimiter(60*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
					ctx.RespondText(rateLimitMsg)
				}),
			},
			&dgc.Command{
				Name:        "balance",
				Description: "returns the current balance of indexed pool tokens held by an account",
				Usage:       " pool balance <pool-name> <account>",
				Example:     " pool balance defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13",
				IgnoreCase:  true,
				Handler:     client.poolBalanceHandler,
				// We want the user to be able to execute this command once in 60 seconds and the cleanup interval shpuld be one second
				RateLimiter: dgc.NewRateLimiter(60*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
					ctx.RespondText(rateLimitMsg)
				}),
			},
			&dgc.Command{
				Name:        "total-value-locked",
				Aliases:     []string{"tvl"},
				Description: "returns the total value locked within a pool in terms of USD. tvl is updated once per hour, providing no <pool-name> parameter returns total value locked across all pools",
				Usage:       " pool total-value-locked <pool-name>",
				Example:     " pool total-value-locked defi5\n!ndx pool total-value-locked",
				IgnoreCase:  true,
				Handler:     client.poolTotalValueLocked,
			},
			&dgc.Command{
				Name:        "total-supply",
				Description: "returns the total supply of pool tokens",
				Usage:       " pool total-supply <pool-name>",
				Example:     " pool total-supply cc10",
				IgnoreCase:  true,
				Handler:     client.poolTotalSupply,
			},
		},
		Usage:   " pool <subcommand> <args...>",
		Example: " pool current-tokens defi5\n!ndx help pool current-tokens",
		Handler: func(ctx *dgc.Ctx) {
			ctx.RespondText("invalid invocation please run a specific subcommand")
		},
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "stake-earned",
		Description: "returns the amount of stake earned. <stake-type> can be one of: defi5, univ2-eth-defi5, cc10, univ2-eth-cc10",
		Usage:       " stake-earned <stake-type> <account>",
		Example:     " stake-earned defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13",
		IgnoreCase:  true,
		Handler:     client.stakeEarnedHandler,
		// We want the user to be able to execute this command once in 60 seconds and the cleanup interval shpuld be one second
		RateLimiter: dgc.NewRateLimiter(60*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText(rateLimitMsg)
		}),
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "uniswap",
		Description: "command group for interacting with Indexed uniswap related contracts",
		Usage:       " uniswap <subcommand> <args...>",
		Example:     " uniswap exchange-amount eth-defi5 1.0\n!ndx uniswap exchange-rate defi5-dai",
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "exchange-amount",
				Description: "returns the amount of tokens received in exchanged for the given pair. \"!ndx uniswap exchange-amount defi5-eth 1.0\" will return the amount of ETH received for swapping 1 DEFI5 to ETH, whereas \"!ndx uniswap exchange-amount eth-defi5 1.0\" will return the amount of DEFI5 received for swapping 1 ETH to DEFI5",
				Usage:       " uniswap exchange-amount <direction> <amount>",
				Example:     " uniswap exchange-amount eth-defi5 1.0\n!ndx uniswap exchange-amount defi5-eth 1.0",
				Handler:     client.uniswapExchangeAmountHandler,
				// We want the user to be able to execute this command once in 60 seconds and the cleanup interval shpuld be one second
				RateLimiter: dgc.NewRateLimiter(60*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
					ctx.RespondText(rateLimitMsg)
				}),
			},
			&dgc.Command{
				Name:        "exchange-rate",
				Description: "returns the exchange rate for a given pair. the <direction> key semantics are the same as with the exchange-amount command",
				Usage:       " uniswap exchange-rate <direction>",
				Example:     " uniswap exchange-rate defi5-dai (returns the value of defi5 in terms of dai)",
				Handler:     client.uniswapExchangeRateHandler,
			},
			&dgc.Command{
				Name:        "price-change",
				Description: "returns the percent price change for a given pair. currently only supports windows in day granularity. the only supported pairs are eth-dai, defi5-dai, cc10-dai, ndx-dai",
				Usage:       " uniswap price-change <pair> <window-in-days>",
				Example:     " uniswap price-change defi5-dai 10 (returns the price change over the last 10 days for defi5)",
				Handler:     client.uniswapPercentChangeHandler,
			},
			&dgc.Command{
				Name:    "price-chart",
				Handler: client.priceWindowChart,
			},
		},

		Handler: func(ctx *dgc.Ctx) {
			ctx.RespondText("invalid invocation please run a specific subcommand")
		},
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "governance",
		Description: "command group for examining the governance contracts",
		Usage:       " governance <subcommand> <args...>",
		Example:     " governance proposal-count\n!ndx governance proposal-info 1",
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "proposal-count",
				Description: "returns the current number of proposals that have been submitted",
				Usage:       " governance proposal-count",
				Example:     " governance proposal-count",
				Handler:     client.governanceCurrentProposalsHandler,
				RateLimiter: dgc.NewRateLimiter(60*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
					ctx.RespondText(rateLimitMsg)
				}),
			},
			&dgc.Command{
				Name:        "proposal-info",
				Usage:       " governance proposal-info <number>",
				Example:     " governance proposal-info 1 (returns information on the first proposal ever submitted)",
				Description: "returns information about the given proposal where <number> represents the proposal submission number. the first proposal submitted would have a submission number of 1, the second proposal submitted would have a submission number of 2, etc...",
				Handler:     client.governanceProposalInfoHandler,
				RateLimiter: dgc.NewRateLimiter(60*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
					ctx.RespondText(rateLimitMsg)
				}),
			},
			&dgc.Command{
				Name:        "total-supply",
				Usage:       " governance total-supply",
				Example:     " governance total-supply",
				Description: "returns the total supply of the governance token, NDX",
				IgnoreCase:  true,
				Handler:     client.governanceTokenTotalSupply,
			},
		},
		Handler: func(ctx *dgc.Ctx) {
			ctx.RespondText("invalid invocation please run a specific subcommand")
		},
	})
	router.Initialize(dg)
	logger.Info("bot started")
	return client, nil
}

// Close terminates the discordgo session
func (c *Client) Close() error {
	c.cancel()
	c.wg.Wait()
	return c.s.Close()
}

func ndxPriceWatchRoutine(ctx context.Context, bot *discordgo.Session, wg *sync.WaitGroup, db *db.Database, logger *zap.Logger) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second * 10)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				defi5TVL, err := db.LastValueLocked("defi5")
				if err != nil {
					logger.Error("failed to get tvl", zap.Error(err), zap.String("asset", "defi5"))
					continue
				}
				cc10TVL, err := db.LastValueLocked("cc10")
				if err != nil {
					logger.Error("failed to get tvl", zap.Error(err), zap.String("asset", "cc10"))
					continue
				}
				orcl5TVL, err := db.LastValueLocked("orcl5")
				if err != nil {
					logger.Error("failed to get tvl", zap.Error(err), zap.String("asset", "orcl5"))
					continue
				}
				degen10TVL, err := db.LastValueLocked("degen10")
				if err != nil {
					logger.Error("failed to get tvl", zap.Error(err), zap.String("asset", "degen10"))
					continue
				}
				nftpTVL, err := db.LastValueLocked("nftp")
				if err != nil {
					logger.Error("failed to get tvl", zap.Error(err), zap.String("asset", "nftp"))
					continue
				}
				errorTVL, err := db.LastValueLocked("error")
				if err != nil {
					logger.Error("fialed to get tvl", zap.Error(err), zap.String("asset", "error"))
				}
				fffTVL, err := db.LastValueLocked("fff")
				if err != nil {
					logger.Error("failed to get tvl", zap.Error(err), zap.String("asset", "fff"))
				}
				price, err := db.LastPrice("ndx")
				if err != nil {
					logger.Error("failed to get price", zap.Error(err), zap.String("asset", "ndx"))
					continue
				}
				guilds, err := bot.UserGuilds(0, "", "")
				if err != nil {
					logger.Error("failed to get user guilds", zap.Error(err))
					continue
				}
				update := fmt.Sprintf("NDXBot: $%0.2f", price)
				// todo(bonedaddy): uncomment when error fund is live
				parsed := ParseValue(defi5TVL + cc10TVL + orcl5TVL + degen10TVL + nftpTVL + errorTVL + fffTVL)
				for _, guild := range guilds {
					bot.GuildMemberNickname(guild.ID, "@me", update)
					bot.UpdateStatus(0, parsed+" TVL")
				}
			}
		}
	}()
}

func launchWatchers(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, bc *bclient.Client, db *db.Database, logger *zap.Logger) error {
	for _, watcher := range cfg.Watchers {
		watcherBot, err := discordgo.New("Bot " + watcher.DiscordToken)
		if err != nil {
			return err
		}
		if err := watcherBot.Open(); err != nil {
			logger.Error("failed to launch watcher bot", zap.Error(err), zap.String("asset", watcher.Currency))
			continue
		}
		switch strings.ToLower(watcher.Currency) {
		case "cc10-defi", "defi5-cc10":
			wg.Add(1)
			go func(name string) {
				defer watcherBot.Close() // register first so it happens before wait gorup closure
				defer wg.Done()
				if err := launchComboWatcherBot(ctx, watcherBot, bc, db, logger, name); err != nil {
					logger.Error("combo watcher bot encountered error", zap.Error(err), zap.String("asset", watcher.Currency))
				}
			}(watcher.Currency)
		default:
			wg.Add(1)
			go func(name string) {
				defer watcherBot.Close() // register first so it happens before wait gorup closure
				defer wg.Done()
				if err := launchSingleWatcherBot(ctx, watcherBot, bc, db, logger, name); err != nil {
					logger.Error("single watcher bot encountered error", zap.Error(err), zap.String("asset", watcher.Currency))
				}
			}(watcher.Currency)
		}
	}
	return nil
}

// used to launch watcher bots that monitor two prices and rotate between each
func launchComboWatcherBot(ctx context.Context, bot *discordgo.Session, bc *bclient.Client, database *db.Database, logger *zap.Logger, name string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		var (
			price float64
			tvl   float64
			err   error
		)
		switch strings.ToLower(name) {
		case "cc10-defi5", "defi5-cc10": // handle combo cc10 and defi5 price tracker
			var defi5DaiPrice, cc10DaiPrice, defi5TVL, cc10TVL float64

			price, err = database.LastPrice("defi5")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "defi5"))
				continue
			}
			tvl, err = database.LastValueLocked("defi5")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "defi5"))
				continue
			}
			defi5TVL = tvl
			defi5DaiPrice = price

			price, err = database.LastPrice("cc10")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "cc10"))
				continue
			}
			tvl, err = database.LastValueLocked("cc10")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "cc10"))
				continue
			}
			cc10TVL = tvl
			cc10DaiPrice = price

			guilds, err := bot.UserGuilds(0, "", "")
			if err != nil {
				logger.Error("failed to get user guilds", zap.Error(err))
				continue
			}

			output := fmt.Sprintf("DEFI5: $%0.2f", defi5DaiPrice)
			parsed := ParseValue(defi5TVL)
			for _, guild := range guilds {
				bot.GuildMemberNickname(guild.ID, "@me", output)
				bot.UpdateStatus(0, parsed+" TVL")
			}

			time.Sleep(time.Second * 10) // wait 10 seconds before displaying new price

			output = fmt.Sprintf("CC10: $%0.2f", cc10DaiPrice)
			parsed = ParseValue(cc10TVL)
			for _, guild := range guilds {
				bot.GuildMemberNickname(guild.ID, "@me", output)
				bot.UpdateStatus(0, parsed+" TVL")
			}

			time.Sleep(time.Second * 10) // wait 10 seconds before looping again
		default:
			return errors.New("unsupported name: " + name)
		}
	}
}

// used to launch watcher bots that monitor one price
func launchSingleWatcherBot(ctx context.Context, bot *discordgo.Session, bc *bclient.Client, database *db.Database, logger *zap.Logger, name string) error {
	ticker := time.NewTicker(time.Second * 30) // only run singles every 30 seconds
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
		var (
			price float64
			tvl   float64
			err   error
		)
		switch strings.ToLower(name) {
		case "defi5":
			price, err = database.LastPrice("defi5")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "defi5"))
				continue
			}
			tvl, err = database.LastValueLocked("defi5")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "defi5"))
				continue
			}
		case "cc10":
			price, err = database.LastPrice("cc10")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "cc10"))
				continue
			}
			tvl, err = database.LastValueLocked("cc10")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "cc10"))
				continue
			}
		case "orcl5":
			price, err = database.LastPrice("orcl5")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "orcl5"))
				continue
			}
			tvl, err = database.LastValueLocked("orcl5")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "orcl5"))
				continue
			}
		case "degen10", "degen":
			// the degen index is named DEGEN but the bot uses the indexed name and its index size
			// internally for referencing pools so override the name
			name = "DEGEN"
			price, err = database.LastPrice("degen10")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "degen10"))
				continue
			}
			tvl, err = database.LastValueLocked("degen10")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "degen10"))
				continue
			}
		case "nftp":
			price, err = database.LastPrice("nftp")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "nftp"))
				continue
			}
			tvl, err = database.LastValueLocked("nftp")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "nftp"))
				continue
			}
		case "error":
			price, err = database.LastPrice("error")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "error"))
				continue
			}
			tvl, err = database.LastValueLocked("error")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "error"))
				continue
			}
		case "fff":
			price, err = database.LastPrice("fff")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "fff"))
				continue
			}
			tvl, err = database.LastValueLocked("fff")
			if err != nil {
				logger.Error("failed to get tvl price", zap.Error(err), zap.String("asset", "fff"))
				continue
			}
		case "ndx":
			price, err = database.LastPrice("ndx")
			if err != nil {
				logger.Error("failed to get dai price", zap.Error(err), zap.String("asset", "ndx"))
				continue
			}
		default:
			return errors.New("unsupported name: " + name)
		}
		guilds, err := bot.UserGuilds(0, "", "")
		if err != nil {
			logger.Error("failed to get user guilds", zap.Error(err))
			continue
		}
		output := printer.Sprintf("%s: $%0.2f", name, price)
		parsed := ParseValue(tvl)
		for _, guild := range guilds {
			bot.GuildMemberNickname(guild.ID, "@me", output)
			bot.UpdateStatus(0, parsed+" TVL")
		}
	}
}
