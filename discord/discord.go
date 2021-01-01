package discord

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
)

var (
	helpEmbed        *discordgo.MessageEmbed
	notifyHelpEmbed  *discordgo.MessageEmbed
	uniswapHelpEmbed *discordgo.MessageEmbed
)

// Client wraps bclient and discordgo to provide a discord bot for indexed finance
type Client struct {
	s  *discordgo.Session
	bc *bclient.Client
	r  *dgc.Router

	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

// NewClient provides a wrapper around discordgo
func NewClient(ctx context.Context, cfg *Config, bc *bclient.Client) (*Client, error) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(ctx)
	for _, watcher := range cfg.Watchers {
		watcherBot, err := discordgo.New("Bot " + watcher.DiscordToken)
		if err != nil {
			cancel()
			return nil, err
		}
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			if err := watcherBot.Open(); err != nil {
				log.Printf("failed to create watcher %s with error: %s", name, err)
				return
			}
			ticker := time.NewTicker(time.Second * 2)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					goto EXIT
				case <-ticker.C:
					break
				}
				var (
					price float64
					err   error
				)
				switch strings.ToLower(name) {
				case "defi5":
					price, err = bc.Defi5DaiPrice()
					if err != nil {
						log.Println("error: ", err)
						goto EXIT
					}
				case "cc10":
					price, err = bc.Cc10DaiPrice()
					if err != nil {
						log.Println("error: ", err)
						goto EXIT
					}
				}
				guilds, err := watcherBot.UserGuilds(0, "", "")
				if err != nil {
					log.Println("error: ", err)
				}
				for _, guild := range guilds {
					watcherBot.GuildMemberNickname(guild.ID, "@me", fmt.Sprintf("%s: $%0.2f", name, price))
				}
				time.Sleep(time.Second * 2)
			}
		EXIT:
			<-ctx.Done()
			watcherBot.Close()
			return
		}(watcher.Currency)
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

	// declare the base router
	router := dgc.Create(&dgc.Router{
		Prefixes:         []string{"!ndx"},
		IgnorePrefixCase: true,  // allow any case of the prefix
		BotsAllowed:      false, // do not allow bots
		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("ping")
		},
	})

	client := &Client{s: dg, bc: bc, r: router, ctx: ctx, cancel: cancel, wg: wg}

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
			},
			&dgc.Command{
				Name:        "balance",
				Description: "returns the current balance of indexed pool tokens held by an account",
				Usage:       " pool balance <pool-name> <account>",
				Example:     " pool balance defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13",
				IgnoreCase:  true,
				Handler:     client.poolBalanceHandler,
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
		Description: "returns the amount of stake earned, currently supports defi5 and univ2-eth-defi5 stake contracts",
		Usage:       " stake-earned <stake-type> <account>",
		Example:     " stake-earned defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13",
		IgnoreCase:  true,
		Handler:     client.stakeEarnedHandler,
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
			},
			&dgc.Command{
				Name:        "exchange-rate",
				Description: "returns the exchange rate for a given pair. the <direction> key semantics are the same as with the exchange-amount command",
				Usage:       " uniswap exchange-rate <direction>",
				Example:     " uniswap exchange-rate defi5-dai (returns the value of defi5 in terms of dai)",
				Handler:     client.uniswapExchangeRateHandler,
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

func (c *Client) sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed); err != nil {
		fmt.Println("error sending message: ", err.Error())
		return
	}
}
