package discord

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/utils"
)

func (c *Client) uniswapExchangeAmountHandler(ctx *dgc.Ctx) {
	if !ctx.Command.RateLimiter.NotifyExecution(ctx) {
		return
	}
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

func (c *Client) uniswapExchangeRateHandler(ctx *dgc.Ctx) {
	arguments := ctx.Arguments
	direction := arguments.Get(0).Raw()
	// valid the allowed currencies
	switch strings.ToLower(direction) {
	case "defi5-dai":
		price, err := c.db.LastPrice("defi5")
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("defi5 dai price fetch failed: ", err)
			return
		}
		ctx.RespondText(fmt.Sprintf("DEFI5-DAI exchange rate: %0.2f", price))
		return
	case "cc10-dai":
		price, err := c.db.LastPrice("cc10")
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("cc10 dai price fetch failed: ", err)
			return
		}
		ctx.RespondText(fmt.Sprintf("CC10-DAI exchange rate: %0.2f", price))
		return
	case "eth-dai":
		price, err := c.db.LastPrice("eth")
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("cc10 dai price fetch failed: ", err)
			return
		}
		ctx.RespondText(fmt.Sprintf("ETH-DAI exchange rate: %0.2f", price))
		return
	case "ndx-dai":
		price, err := c.db.LastPrice("ndx")
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("ndx dai price fetch failed: ", err)
			return
		}
		ctx.RespondText(fmt.Sprintf("NDX-DAI exchange rate: %0.2f", price))
	default:
		ctx.RespondText("invalid currency requested must be one of: defi5-dai, cc10-dai, eth-dai, ndx-dai")
		return
	}
}

func (c *Client) uniswapPercentChangeHandler(ctx *dgc.Ctx) {
	arguments := ctx.Arguments
	pair := arguments.Get(0).Raw()
	window, err := arguments.Get(1).AsInt()
	if err != nil {
		ctx.RespondText("window argument is not a number")
		return
	}
	// valid the allowed currencies
	switch strings.ToLower(pair) {
	case "defi5-dai":
		price, err := c.db.PriceChangeInRange("defi5", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("defi5 dai price change fetch failed: ", err)
			return
		}
		var changed string
		if (price * 100) < 0 {
			changed = "decreased"
		} else {
			changed = "increased"
		}
		ctx.RespondText(fmt.Sprintf("DEFI5-DAI price has %s %0.2f%% over the last %v days", changed, math.Abs(price*100), window))
		return
	case "cc10-dai":
		price, err := c.db.PriceChangeInRange("cc10", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("cc10 dai price change fetch failed: ", err)
			return
		}
		var changed string
		if (price * 100) < 0 {
			changed = "decreased"
		} else {
			changed = "increased"
		}
		ctx.RespondText(fmt.Sprintf("CC10-DAI price has %s %0.2f%% over the last %v days", changed, math.Abs(price*100), window))
		return
	case "eth-dai":
		price, err := c.db.PriceChangeInRange("eth", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("cc10 dai price change failed: ", err)
			return
		}
		var changed string
		if (price * 100) < 0 {
			changed = "decreased"
		} else {
			changed = "increased"
		}
		ctx.RespondText(fmt.Sprintf("ETH-DAI price has %s %0.2f%% over the last %v days", changed, math.Abs(price*100), window))
		return
	case "ndx-dai":
		price, err := c.db.PriceChangeInRange("ndx", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("ndx dai price change failed: ", err)
			return
		}
		var changed string
		if (price * 100) < 0 {
			changed = "decreased"
		} else {
			changed = "increased"
		}
		ctx.RespondText(fmt.Sprintf("NDX-DAI price has %s %0.2f%% over the last %v days", changed, math.Abs(price*100), window))
	default:
		ctx.RespondText("invalid currency requested must be one of: defi5-dai, cc10-dai, eth-dai, ndx-dai")
		return
	}
}
