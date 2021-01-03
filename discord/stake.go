package discord

import (
	"fmt"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) stakeEarnedHandler(ctx *dgc.Ctx) {
	if !ctx.Command.RateLimiter.NotifyExecution(ctx) {
		return

	}
	arguments := ctx.Arguments
	stakeType := arguments.Get(0).Raw()
	accountAddr := arguments.Get(1).Raw()
	var poolType string
	switch stakeType {
	case "univ2-eth-defi5", "defi5":
		poolType = "defi5"
	default:
		ctx.RespondText("invalid stake-type")
		return
	}
	ip, err := c.getIndexPool(poolType)
	if err != nil {
		ctx.RespondText("failed to lookup index pool")
		return
	}
	sp, err := c.getStakingRewards(stakeType)
	if err != nil {
		ctx.RespondText("invalid stake-type")
		return
	}
	earned, err := bclient.StakeEarned(sp, ip, common.HexToAddress(accountAddr))
	if err != nil {
		ctx.RespondText("failed to lookup staking balance")
		return
	}
	earnedF, _ := earned.Float64()
	ctx.RespondText(fmt.Sprintf("staking rewards earned for %s: %0.2f", accountAddr, earnedF))
}
