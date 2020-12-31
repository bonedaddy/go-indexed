package discord

import (
	"fmt"
	"log"

	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

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
		sp, err = c.bc.StakingAt(args[2])
		if err != nil {
			break
		}
		ip, err = c.bc.DEFI5()
	case "univ2-eth-defi5":
		sp, err = c.bc.StakingAt(args[2])
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
