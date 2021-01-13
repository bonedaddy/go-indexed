package discord

import (
	"fmt"
	"log"
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) poolTokensHandler(ctx *dgc.Ctx) {
	if !ctx.Command.RateLimiter.NotifyExecution(ctx) {
		return

	}
	arguments := ctx.Arguments
	poolName := arguments.Get(0).Raw()
	ip, err := c.getIndexPool(poolName)
	if err != nil {
		ctx.RespondText("invalid pool")
		return
	}
	tokens, err := c.bc.PoolTokensFor(ip)
	if err != nil {
		ctx.RespondText("failed to lookup current tokens")
		return
	}
	tokensEmbed := BaseEmbed()
	tokensEmbed.Title = fmt.Sprintf("%s Current Pool Tokens", strings.ToUpper(poolName))

	for name, addr := range tokens {
		tokensEmbed.Fields = append(tokensEmbed.Fields, &discordgo.MessageEmbedField{
			Name:  name,
			Value: addr.String(),
		})
	}
	ctx.RespondEmbed(tokensEmbed)
}

func (c *Client) poolBalanceHandler(ctx *dgc.Ctx) {
	if !ctx.Command.RateLimiter.NotifyExecution(ctx) {
		return
	}
	arguments := ctx.Arguments
	poolName := arguments.Get(0).Raw()
	accountAddr := arguments.Get(1).Raw()
	ip, err := c.getIndexPool(poolName)
	if err != nil {
		ctx.RespondText("invalid pool")
		return
	}
	bal, err := bclient.BalanceOfDecimal(ip, common.HexToAddress(accountAddr))
	if err != nil {
		ctx.RespondText("failed to lookup balance")
		return
	}
	balF, _ := bal.Float64()
	ctx.RespondText(fmt.Sprintf("account balance for %s: %0.2f", accountAddr, balF))
}

func (c *Client) poolTotalValueLocked(ctx *dgc.Ctx) {
	if !ctx.Command.RateLimiter.NotifyExecution(ctx) {
		return
	}
	arguments := ctx.Arguments
	poolName := arguments.Get(0).Raw()
	ip, err := c.getIndexPool(poolName)
	if err != nil {
		ctx.RespondText("invalid pool")
		return
	}
	tvl, err := c.bc.GetTotalValueLocked(ip)
	if err != nil {
		log.Println("failed to get total value locked: ", err)
		ctx.RespondText("failed to get total value locked")
		return
	}
	ctx.RespondText(fmt.Sprintf("total value locked for %s: %0.2f", poolName, tvl))
}
