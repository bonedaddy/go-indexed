package discord

import (
	"fmt"
	"log"
	"strings"

	"github.com/bonedaddy/go-indexed/bclient"
	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

var (
	helpEmbed *discordgo.MessageEmbed
)

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

func (c *Client) handleCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 && args[0] == "!ndx" {
		c.sendHelp(s, m)
		return
	}
	if len(args) > 2 {
		switch args[1] {
		case "pool-balance":
			c.poolBalance(s, m, args)
		case "stake-earned":
			c.stakeEarned(s, m, args)
		}
	}
}

func (c *Client) stakeEarned(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1              2            3
	// !ndx staking-earned <stake-type> <account>
	var (
		sp  *stakingbindings.Stakingbindings
		ip  bclient.IndexPool
		err error
	)
	switch args[2] {
	case "defi5":
		sp, err = c.bc.StakingAt("defi5")
		if err != nil {
			break
		}
		ip, err = c.bc.DEFI5()
	default:
		_, err = c.s.ChannelMessageSend(m.ChannelID, "invalid stake-type")
		if err != nil {
			log.Println("failed to send channel message")
		}
		return
	}
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "failed to lookup staking balance")
		log.Println("error encountered: ", err)
		return
	}
	earned, err := bclient.StakeEarned(sp, ip, common.HexToAddress(args[3]))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "failed to lookup staking balance")
		return
	}
	c.s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("staking rewards earned for %s: %s", args[3], earned))
}

func (c *Client) poolBalance(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1            2      3
	// !ndx pool-balance <pool> <account>
	var (
		ip  bclient.IndexPool
		err error
	)
	switch args[2] {
	case "defi5":
		ip, err = c.bc.DEFI5()
	case "cc10":
		ip, err = c.bc.CC10()
	default:
		_, err = c.s.ChannelMessageSend(m.ChannelID, "invalid pool")
		if err != nil {
			log.Println("failed to send channel message")
		}
		return
	}
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "failed to get index pool binding")
		log.Println("failed to get index pool binding")
		return
	}
	bal, err := bclient.BalanceOfDecimal(ip, common.HexToAddress(args[3]))
	if err != nil {
		_, err = c.s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("failed to lookup balance: %s", err))
		if err != nil {
			log.Println("failed to send channel message")
		}
	}
	c.s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("account balance for %s: %s", args[3], bal))
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
