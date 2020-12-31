package discord

import (
	"fmt"
	"strconv"

	"github.com/bonedaddy/go-indexed/utils"
	"github.com/bwmarrin/discordgo"
)

func (c *Client) handleUniswap(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// 0     1      2
	// !ndx uniswap <command> <args...>
	/*
		where command and args can be:
			exchange-amount <direction> <amount>
				-> exchange-amount eth-defi5 1 (exchange amount swapping 1 ETH for DEFI5)
				-> exchange amount defi5-eth 10 (exchange amount swapping 10 DEFI5 for ETH)
	*/
	switch args[2] {
	case "help":
		c.s.ChannelMessageSendEmbed(m.ChannelID, uniswapHelpEmbed)
		return
	case "exchange-amount":
		// 0    1       2               3      4
		// !ndx uniswap exchange-amount <pair> <amount-float>
		if len(args) < 5 {
			c.s.ChannelMessageSend(m.ChannelID, "invalid invocation of !ndx uniswap exchange-amount")
			return
		}
		amtF, err := strconv.ParseFloat(args[4], 64)
		if err != nil {
			c.s.ChannelMessageSend(m.ChannelID, "invalid invocation of !ndx uniswap exchange-amount")
			return
		}
		amt := utils.ToWei(amtF, c.bc.PairDecimals(args[3]))
		exchAmt, err := c.bc.ExchangeAmount(amt, args[3])
		if err != nil {
			c.s.ChannelMessageSend(m.ChannelID, "failed to get exchange amount: "+err.Error())
			return
		}
		c.s.ChannelMessageSend(m.ChannelID,
			fmt.Sprintf("swapping %v with pair %s will yield: %s", amtF, args[3], utils.ToDecimal(exchAmt, c.bc.PairDecimals(args[3]))),
		)
	}
}
