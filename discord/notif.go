package discord

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) handleNotif(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1      2
	// !ndx notify <command> <args...> [-interval]
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
			var newArgs []string
			newArgs = append(newArgs,
				args[0], "staking-earned", args[3], args[4],
			)
			c.stakeEarned(s, m, newArgs)
			time.Sleep(interval)
			count++
		}
	}
}
