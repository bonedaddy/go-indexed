package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) handleNotif(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0    1      2
	// !ndx notify <command> <args...>
	for {
		// temporary
		switch args[2] {
		case "stake-earned":
			// 0    1      2             3           4
			// !ndx notify stake-balance <pool-name> <address>
			if len(args) != 5 {
				c.s.ChannelMessageSend(m.ChannelID, "invalid invocation of notify stake-balance command")
				return
			}
			var newArgs []string
			newArgs = append(newArgs,
				args[0], "staking-earned", args[3], args[4],
			)
			c.stakeEarned(s, m, newArgs)
			time.Sleep(time.Second * 60)
		}
	}
}
