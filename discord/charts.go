package discord

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/db"
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
	var (
		prices []*db.Price
	)
	// valid the allowed currencies
	switch strings.ToLower(pair) {
	case "defi5-dai":
		prices, err = c.db.PricesInRange("defi5", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("defi5 dai price change fetch failed: ", err)
			return
		}
	case "cc10-dai":
		prices, err = c.db.PricesInRange("cc10", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("cc10 dai price change fetch failed: ", err)
			return
		}
	case "ndx-dai":
		prices, err = c.db.PricesInRange("ndx", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("ndx dai price change fetch failed: ", err)
			return
		}
	case "eth-dai":
		prices, err = c.db.PricesInRange("eth", window)
		if err != nil {
			ctx.RespondText("failed to get price")
			log.Println("eth dai price change fetch failed: ", err)
			return
		}
	default:
		ctx.RespondText("invalid currency requested must be one of: defi5-dai, cc10-dai, eth-dai, ndx-dai")
		return
	}
	mainSeries := chart.TimeSeries{}
	for _, price := range prices {
		mainSeries.XValues = append(mainSeries.XValues, price.CreatedAt)
		// mainSeries.XValues = append(mainSeries.XValues, float64(i))
		mainSeries.YValues = append(mainSeries.YValues, price.USDPrice)
	}

	smaSeries := &chart.SMASeries{
		InnerSeries: mainSeries,
	}

	minSeries := &chart.MinSeries{
		Name: "sma min",
		Style: chart.Style{
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: smaSeries,
	}

	maxSeries := &chart.MaxSeries{
		Name: "sma max",
		Style: chart.Style{
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: smaSeries,
	}

	graph := chart.Chart{
		Title:  fmt.Sprintf("%s %v day window", pair, window),
		Width:  1920,
		Height: 1080,
		Series: []chart.Series{
			mainSeries,
			smaSeries,
			minSeries,
			maxSeries,
		},
		XAxis: chart.XAxis{
			// ensure we render date timestamps with minute granularity
			ValueFormatter: chart.TimeMinuteValueFormatter,
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
}
