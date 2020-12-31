package discord

import (
	"fmt"
	"log"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) stakeEarned(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1              2            3
	// !ndx staking-earned <stake-type> <account>
	ip, err := c.getIndexPool(args[2])
	if err != nil {
		_, err = c.s.ChannelMessageSend(m.ChannelID, "invalid stake-type")
		if err != nil {
			log.Println("failed to send channel message")
		}
	}
	sp, err := c.getStakingRewards(args[2])
	if err != nil {
		_, err = c.s.ChannelMessageSend(m.ChannelID, "invalid stake-type")
		if err != nil {
			log.Println("failed to send channel message")
		}
	}
	earned, err := bclient.StakeEarned(sp, ip, common.HexToAddress(args[3]))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "failed to lookup staking balance")
		return
	}
	c.s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("staking rewards earned for %s: %s", args[3], earned))
}
