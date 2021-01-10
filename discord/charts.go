package discord

import (
	"bytes"
	"log"
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/wcharczuk/go-chart/v2"
)

func (c *Client) priceWindowChart(ctx *dgc.Ctx) {

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

		prices, err := c.db.PricesInRange("defi5", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("defi5 dai price change fetch failed: ", err)
			return
		}
		mainSeries := chart.ContinuousSeries{
			Name: "DEFI5-DAI SMA",
			// XValues are unix timestamps
			// YValues are prices
		}
		for _, price := range prices {
			mainSeries.XValues = append(mainSeries.XValues, float64(price.CreatedAt.Unix()))
			mainSeries.YValues = append(mainSeries.YValues, price.USDPrice)
		}

		smaSeries := &chart.SMASeries{
			InnerSeries: mainSeries,
		}
		graph := chart.Chart{
			Series: []chart.Series{
				mainSeries,
				smaSeries,
			},
		}
		buffer := bytes.NewBuffer(nil)
		if err := graph.Render(chart.PNG, buffer); err != nil {
			log.Println("failed to render SMA: ", err)
			ctx.RespondText("failed to render SMA")
			return
		}
		if _, err := ctx.Session.ChannelFileSend(ctx.Event.ChannelID, "chart.png", buffer); err != nil {
			log.Println("failed to upload chart: ", err)
			ctx.RespondText("failed to upload chart")
			return
		}
		return
	default:
		ctx.RespondText("invalid currency requested must be one of: defi5-dai, cc10-dai, eth-dai, ndx-dai")
		return
	}
}
