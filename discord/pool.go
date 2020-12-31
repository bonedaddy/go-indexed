package discord

import (
	"fmt"
	"log"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) poolTokens(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1           2
	// !ndx pool-tokens <pool-name>
	ip, err := c.getIndexPool(args[2])
	if err != nil {
		_, err = c.s.ChannelMessageSend(m.ChannelID, "invalid pool")
		if err != nil {
			log.Println("failed to send channel message")
		}
		return
	}
	tokens, err := c.bc.PoolTokensFor(ip)
	if err != nil {
		_, err = c.s.ChannelMessageSend(m.ChannelID, "failed to lookup current tokens")
		if err != nil {
			log.Println("failed to send channel message")
		}
		return
	}
	msg := fmt.Sprintf("current tokens in pool %s\n", args[2])
	for name, addr := range tokens {
		msg += fmt.Sprintf("%s (%s)\n", name, addr)
	}
	c.s.ChannelMessageSend(m.ChannelID, msg)
}

func (c *Client) poolBalance(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1            2      3
	// !ndx pool-balance <pool> <account>
	ip, err := c.getIndexPool(args[2])
	if err != nil {
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
