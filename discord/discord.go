package discord

import (
	"fmt"
	"log"
	"strings"

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
}

// NewClient provides a wrapper around discordgo
func NewClient(token string, bc *bclient.Client) (*Client, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	client := &Client{token: token, s: dg, bc: bc}
	client.s.AddHandler(client.messageCreate)
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
	if len(args) > 2 {
		switch args[1] {
		case "pool-tokens":
			c.poolTokens(s, m, args)
			return
		}
	}
	if len(args) > 3 {
		switch args[1] {
		case "pool-balance":
			c.poolBalance(s, m, args)
			return
		case "stake-earned":
			c.stakeEarned(s, m, args)
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
