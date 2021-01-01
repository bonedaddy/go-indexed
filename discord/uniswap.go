package discord

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/utils"
)

func (c *Client) uniswapExchangeAmountHandler(ctx *dgc.Ctx) {
	arguments := ctx.Arguments
	direction := arguments.Get(0).Raw()
	amtF, err := strconv.ParseFloat(arguments.Get(1).Raw(), 64)
	if err != nil {
		ctx.RespondText("invalid amount given")
		return
	}
	amt := utils.ToWei(amtF, c.bc.PairDecimals(direction))
	exchAmt, err := c.bc.ExchangeAmount(amt, direction)
	if err != nil {
		ctx.RespondText("failed to get exchange amount")
		return
	}
	exchDec := utils.ToDecimal(exchAmt, c.bc.PairDecimals(direction))
	exchF, _ := exchDec.Float64()
	pairParts := strings.Split(direction, "-")
	if len(pairParts) < 2 {
		ctx.RespondText("failed to parse given pair")
		return
	}
	inPair := strings.ToUpper(pairParts[0])
	outPair := strings.ToUpper(pairParts[1])
	ctx.RespondText(
		fmt.Sprintf("swapping %v-%s will yield approximately %0.2f-%s\nnote: all prices are approximates", amtF, inPair, exchF, outPair),
	)
}
