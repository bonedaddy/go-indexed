package discord

import (
	"fmt"
	"log"
	"strings"

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
	// declare the base router
	router := dgc.Create(&dgc.Router{
		Prefixes:         []string{"!ndx"},
		IgnorePrefixCase: true,  // allow any case of the prefix
		BotsAllowed:      false, // do not allow bots
	})

	client := &Client{token: token, s: dg, bc: bc, r: router}

	router.RegisterDefaultHelpCommand(dg, nil)

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

	router.Initialize(dg)

	// client.s.AddHandler(client.messageCreate)
	if err := dg.Open(); err != nil {
		return nil, err
	}
	log.Println("bot is now running")
	return client, nil
}

// Close terminates the discordgo session
func (c *Client) Close() error {
	return c.s.Close()
}

// handleCommand processes incoming messages from discord and determines what to do with them
func (c *Client) handleCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 && args[0] == "!ndx" {
		c.sendHelp(s, m)
		return
	} else if len(args) == 2 && args[0] == "!ndx" && args[1] == "help" {
		c.sendHelp(s, m)
		return
	}

	if len(args) >= 2 {
		if args[1] == "notify" {
			c.handleNotif(s, m, args)
			return
		}
		if args[1] == "uniswap" {
			c.handleUniswap(s, m, args)
			return
		}
	}
	c.s.ChannelMessageSend(m.ChannelID, "invalid command invocation")
}

func (c *Client) sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed); err != nil {
		fmt.Println("error sending message: ", err.Error())
		return
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func (c *Client) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// parse the message contents based off string fields
	args := strings.Fields(m.Content)
	if len(args) == 0 {
		return
	}
	// ensure the first field is a valid invocation of index bot1
	if args[0] != "!ndx" {
		fmt.Println("invalid invocation")
		return
	}
	c.handleCommand(s, m, args)
}
