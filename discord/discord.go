package discord

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	helpEmbed *discordgo.MessageEmbed
)

type Client struct {
	token string
	s     *discordgo.Session
}

func init() {
	helpEmbed = &discordgo.MessageEmbed{
		Title:       "Indexed Finance Bot Help Menu",
		Description: "All commands must be invoked with !ndx <cmd>\nAnything with <..> after command name expects an argument\nAnything with [..] after command name is an optional argument",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://pbs.twimg.com/profile_images/1342395531318976518/kIv5abLc_400x400.jpg",
		},
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "pool-balance <pool-name> <account-address>",
				Value: "returns the index pool balance for a given pool held by account-address",
			},
		},
	}
}

// NewClient provides a wrapper around discordgo
func NewClient(token string) (*Client, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	dg.AddHandler(messageCreate)
	if err := dg.Open(); err != nil {
		return nil, err
	}
	log.Println("bot is now running")
	return &Client{token, dg}, nil
}

// Close terminates the discordgo session
func (c *Client) Close() error {
	return c.s.Close()
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 && args[0] == "!ndx" {
		sendHelp(s, m)
		return
	}
}

func sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed); err != nil {
		fmt.Println("error sending message: ", err.Error())
		return
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	handleCommand(s, m, args)
}
