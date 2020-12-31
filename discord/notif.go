package discord

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) handleNotif(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1      2
	// !ndx notify <command> <args...> [-interval]

	if args[2] == "help" {
		c.s.ChannelMessageSendEmbed(m.ChannelID, notifyHelpEmbed)
		return
	}

	var interval = time.Second * 60
	if strings.Contains(args[len(args)-1], "-interval") {
		parts := strings.Split(args[len(args)-1], "=")
		if len(parts) == 2 {
			i, err := strconv.Atoi(parts[1])
			if err != nil {
				c.s.ChannelMessageSend(m.ChannelID, "invalid -interval option")
				return
			}
			interval = time.Second * time.Duration(i)
		}

	}
	count := 0
	for {
		if count+1 == 10 {
			break
		}
		// temporary
		switch args[2] {
		case "stake-earned":
			// 0    1      2             3           4
			// !ndx notify stake-balance <pool-name> <address>
			if len(args) < 5 {
				c.s.ChannelMessageSend(m.ChannelID, "invalid invocation of notify stake-balance command")
				return
			}
			ip, err := c.getIndexPool(args[3])
			if err != nil {
				c.s.ChannelMessageSend(m.ChannelID, err.Error())
				return
			}
			sp, err := c.getStakingRewards(args[3])
			if err != nil {
				c.s.ChannelMessageSend(m.ChannelID, err.Error())
				return
			}
			earned, err := bclient.StakeEarned(sp, ip, common.HexToAddress(args[4]))
			if err != nil {
				c.s.ChannelMessageSend(m.ChannelID, "failed to get stake earned")
				return
			}
			_, err = c.s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s: staking rewards earned for pool %s: %s", m.Author.Mention(), args[3], earned))
			if err != nil {
				log.Println("failed to send message: ", err)
				return
			}
			time.Sleep(interval)
			count++
		}
	}
}
