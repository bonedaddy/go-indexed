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
		graph := chart.Chart{
			Title: "DEFI5-DAI Bar Chart",
			Background: chart.Style{
				Padding: chart.Box{
					Top: 40,
				},
			},
			Height: 512,
			Series: []chart.Series{
				chart.ContinuousSeries{
					// stuff here
				},
			},
		}
		for _, price := range prices {
			graph.Bars = append(graph.Bars, chart.Value{
				Label: price.CreatedAt.String(),
				Value: price.USDPrice,
			})
		}
		buffer := bytes.NewBuffer(nil)
		if err := graph.Render(chart.PNG, buffer); err != nil {
			log.Println("error rendering: ", err)
			return
		}
		ctx.Session.ChannelFileSend(ctx.Event.ChannelID, "chart.png", buffer)
		return
	default:
		ctx.RespondText("invalid currency requested must be one of: defi5-dai, cc10-dai, eth-dai, ndx-dai")
		return
	}
}
