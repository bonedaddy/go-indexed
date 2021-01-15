package discord

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/db"
	"github.com/bwmarrin/discordgo"
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
}

func init() {
	printer = message.NewPrinter(message.MatchLanguage("en"))
}

// NewClient provides a wrapper around discordgo
func NewClient(ctx context.Context, cfg *Config, bc *bclient.Client, db *db.Database) (*Client, error) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(ctx)
	// launch price watcher bots
	if err := launchWatchers(ctx, wg, cfg, bc, db); err != nil {
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
	if err := dg.UpdateListeningStatus("!ndx help"); err != nil {
		log.Println("failed to udpate streaming status: ", err)
	}

	// updates the main bot's nickname to reflect ndx price
	ndxPriceWatchRoutine(ctx, dg, wg, db)

	// declare the base router
	router := dgc.Create(&dgc.Router{
		Prefixes:         []string{"!ndx"},
		IgnorePrefixCase: true,  // allow any case of the prefix
		BotsAllowed:      false, // do not allow bots
		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("ping")
		},
	})

	client := &Client{s: dg, bc: bc, r: router, ctx: ctx, cancel: cancel, wg: wg, db: db}

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
				Description: "returns the total value locked within a pool in terms of USD. tvl is updated once per hour",
				Usage:       " pool total-value-locked <pool-name>",
				Example:     " pool total-value-locked defi5",
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

	log.Println("bot is now running")
	return client, nil
}

// Close terminates the discordgo session
func (c *Client) Close() error {
	c.cancel()
	c.wg.Wait()
	return c.s.Close()
}

func ndxPriceWatchRoutine(ctx context.Context, bot *discordgo.Session, wg *sync.WaitGroup, db *db.Database) {
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
				price, err := db.LastPrice("ndx")
				if err != nil {
					log.Println("failed to get last ndx price: ", err)
					continue
				}
				guilds, err := bot.UserGuilds(0, "", "")
				if err != nil {
					log.Println("failed to get user guilds error: ", err)
					continue
				}
				for _, guild := range guilds {
					bot.GuildMemberNickname(guild.ID, "@me", fmt.Sprintf("NDXBot: $%0.2f", price))
				}
			}
		}
	}()
}

func launchWatchers(ctx context.Context, wg *sync.WaitGroup, cfg *Config, bc *bclient.Client, db *db.Database) error {
	for _, watcher := range cfg.Watchers {
		watcherBot, err := discordgo.New("Bot " + watcher.DiscordToken)
		if err != nil {
			return err
		}
		if err := watcherBot.Open(); err != nil {
			log.Printf("failed to launch watcher bot %s: %s", watcher.Currency, err)
			continue
		}
		// set playing status
		watcherBot.UpdateStatus(0, "indexed.finance")
		switch strings.ToLower(watcher.Currency) {
		case "cc10-defi", "defi5-cc10":
			wg.Add(1)
			go func(name string) {
				defer watcherBot.Close() // register first so it happens before wait gorup closure
				defer wg.Done()
				if err := launchComboWatcherBot(ctx, watcherBot, bc, db, name); err != nil {
					log.Printf("an error occured for %s price watcher: %s", name, err)
				}
			}(watcher.Currency)
		default:
			wg.Add(1)
			go func(name string) {
				defer watcherBot.Close() // register first so it happens before wait gorup closure
				defer wg.Done()
				if err := launchSingleWatcherBot(ctx, watcherBot, bc, db, name); err != nil {
					log.Printf("an error occured for %s price watcher: %s", name, err)
				}
			}(watcher.Currency)
		}
	}
	return nil
}

// used to launch watcher bots that monitor two prices and rotate between each
func launchComboWatcherBot(ctx context.Context, bot *discordgo.Session, bc *bclient.Client, database *db.Database, name string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		var (
			price float64
			err   error
		)
		switch strings.ToLower(name) {
		case "cc10-defi5", "defi5-cc10": // handle combo cc10 and defi5 price tracker
			var defi5DaiPrice, cc10DaiPrice float64

			price, err = database.LastPrice("defi5")
			if err != nil {
				log.Println("failed to get defi5 dai price: ", err)
				continue
			}
			defi5DaiPrice = price

			price, err = database.LastPrice("cc10")
			if err != nil {
				log.Println("failed to get cc10 dai price: ", err)
				continue
			}
			cc10DaiPrice = price

			guilds, err := bot.UserGuilds(0, "", "")
			if err != nil {
				log.Println("failed to get user guilds error: ", err)
				continue
			}
			for _, guild := range guilds {
				bot.GuildMemberNickname(guild.ID, "@me", fmt.Sprintf("DEFI5: $%0.2f", defi5DaiPrice))
			}
			time.Sleep(time.Second * 10) // wait 10 seconds before displaying new price
			for _, guild := range guilds {
				bot.GuildMemberNickname(guild.ID, "@me", fmt.Sprintf("CC10: $%0.2f", cc10DaiPrice))
			}
			time.Sleep(time.Second * 10) // wait 10 seconds before looping again
		default:
			return errors.New("unsupported name: " + name)
		}
	}
}

// used to launch watcher bots that monitor one price
func launchSingleWatcherBot(ctx context.Context, bot *discordgo.Session, bc *bclient.Client, database *db.Database, name string) error {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
		var (
			price float64
			err   error
		)
		switch strings.ToLower(name) {
		case "defi5":
			price, err = database.LastPrice("defi5")
			if err != nil {
				log.Println("failed to get defi5 dai price error: ", err)
				continue
			}
		case "cc10":
			price, err = database.LastPrice("cc10")
			if err != nil {
				log.Println("failed to get cc10 dai price error: ", err)
				continue
			}
		case "ndx":
			price, err = database.LastPrice("ndx")
			if err != nil {
				log.Println("failed to get ndx dai price error: ", err)
				continue
			}
		default:
			return errors.New("unsupported name: " + name)
		}
		guilds, err := bot.UserGuilds(0, "", "")
		if err != nil {
			log.Println("failed to get user guilds error: ", err)
			continue
		}
		for _, guild := range guilds {
			bot.GuildMemberNickname(guild.ID, "@me", fmt.Sprintf("%s: $%0.2f", name, price))
		}
	}
}
