package discord

import (
	"fmt"
	"log"

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
	token string
	s     *discordgo.Session
	bc    *bclient.Client
	r     *dgc.Router
}

// NewClient provides a wrapper around discordgo
func NewClient(token string, bc *bclient.Client) (*Client, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	if err := dg.Open(); err != nil {
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

	client := &Client{token: token, s: dg, bc: bc, r: router}

	// register our custom help command
	registerHelpCommand(dg, nil, router)

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
		Example:     " uniswap exchange-amount eth-defi5 1.0",
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "exchange-amount",
				Description: "returns the amount of tokens received in exchanged for the given pair. \"!ndx uniswap exchange-amount defi5-eth 1.0\" will return the amount of ETH received for swapping 1 DEFI5 to ETH, whereas \"!ndx uniswap exchange-amount eth-defi5 1.0\" will return the amount of DEFI5 received for swapping 1 ETH to DEFI5",
				Usage:       " uniswap exchange-amount <direction> <amount>",
				Example:     " uniswap exchange-amount eth-defi5 1.0\n!ndx uniswap exchange-amount defi5-eth 1.0",
				Handler:     client.uniswapExchangeAmountHandler,
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
	return c.s.Close()
}

func (c *Client) sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed); err != nil {
		fmt.Println("error sending message: ", err.Error())
		return
	}
}
