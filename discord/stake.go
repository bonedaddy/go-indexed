package discord

import (
	"fmt"
	"log"
	"strings"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) stakeEarned(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1            2            3
	// !ndx stake-earned <stake-type> <account>
	var poolName string
	if strings.Contains(args[2], "defi5") {
		poolName = "defi5"
	} else {
		c.s.ChannelMessageSend(m.ChannelID, "invalid stake type")
		return
	}
	ip, err := c.getIndexPool(poolName)
	if err != nil {
		c.s.ChannelMessageSend(m.ChannelID, "invalid stake-type")
		log.Println("getIndexPool error: ", err)
		return
	}
	sp, err := c.getStakingRewards(args[2])
	if err != nil {
		c.s.ChannelMessageSend(m.ChannelID, "invalid stake-type")
		log.Println("getStakingRewards error: ", err)
		return
	}
	earned, err := bclient.StakeEarned(sp, ip, common.HexToAddress(args[3]))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "failed to lookup staking balance")
		return
	}
	earnedF, _ := earned.Float64()
	c.s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("staking rewards earned for %s: %0.2f", args[3], earnedF))
}
