package discord

import (
	"fmt"
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

func (c *Client) poolTokensHandler(ctx *dgc.Ctx) {
	if !ctx.Command.RateLimiter.NotifyExecution(ctx) {
		return

	}
	arguments := ctx.Arguments
	poolName := arguments.Get(0).Raw()
	ip, err := c.bc.GetIndexPool(poolName)
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
	ip, err := c.bc.GetIndexPool(poolName)
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
	arguments := ctx.Arguments
	if arguments.Amount() == 0 {
		// get tvl across all pools
		var (
			totalLocked float64
			pools       = []string{"defi5", "cc10", "orcl5", "degen10"}
		)
		for _, pool := range pools {
			tvl, err := c.db.LastValueLocked(strings.ToLower(pool))
			if err != nil {
				c.logger.Error("failed to get total value locked", zap.Error(err), zap.String("asset", pool))
				continue // instead of bailing log the error and continue, this makes it easier to roll out price tracking updates sooner
			}
			totalLocked += tvl
		}
		ctx.RespondText(printer.Sprintf("total value locked across all pools: $%0.2f", totalLocked))
	} else {
		poolName := arguments.Get(0).Raw()
		// handle cases where degen10 is given as degen
		if strings.ToLower(poolName) == "degen" {
			poolName = "degen10"
		}
		tvl, err := c.db.LastValueLocked(strings.ToLower(poolName))
		if err != nil {
			c.logger.Error("failed to get total value locked", zap.Error(err), zap.String("asset", poolName))
			ctx.RespondText("failed to get total value locked")
			return
		}
		ctx.RespondText(printer.Sprintf("total value locked for %s: $%0.2f", poolName, tvl))
	}
}

func (c *Client) poolTotalSupply(ctx *dgc.Ctx) {
	arguments := ctx.Arguments
	poolName := arguments.Get(0).Raw()
	supply, err := c.db.LastTotalSupply(poolName)
	if err != nil {
		ctx.RespondText("invalid pool")
		return
	}
	ctx.RespondText(printer.Sprintf("total supply for %s: %0.2f", poolName, supply))
}
