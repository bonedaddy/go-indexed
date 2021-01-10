package discord

import (
	"bytes"
	"log"
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/db"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
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
	priceSeries := chart.TimeSeries{
		Name: pair,
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
	}
	// records prices at every hour
	priceHourSeries := chart.TimeSeries{
		Name: pair,
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
	}
	var lastHour int
	for _, price := range prices {
		if price.CreatedAt.Hour() == 0 {
			// reset
			lastHour = 0
		}
		if lastHour == 0 {
			lastHour = price.CreatedAt.Minute()
			priceHourSeries.XValues = append(priceHourSeries.XValues, price.CreatedAt)
			priceHourSeries.YValues = append(priceHourSeries.YValues, price.USDPrice)
		}
		if price.CreatedAt.Hour() > lastHour {
			lastHour = price.CreatedAt.Hour()
			priceHourSeries.XValues = append(priceHourSeries.XValues, price.CreatedAt)
			priceHourSeries.YValues = append(priceHourSeries.YValues, price.USDPrice)
		}
		priceSeries.XValues = append(priceSeries.XValues, price.CreatedAt)
		priceSeries.YValues = append(priceSeries.YValues, price.USDPrice)
	}

	hourlySMASeries := &chart.SMASeries{
		Name:   pair + " - " + "sma",
		Period: window,
		Style: chart.Style{
			StrokeColor: drawing.ColorRed,
		},
		InnerSeries: priceHourSeries,
	}

	bbSeries := &chart.BollingerBandsSeries{
		Name:   pair + " - bol. bands",
		Period: window,
		K:      2,
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("efefef"),
			FillColor:   drawing.ColorFromHex("efefef").WithAlpha(64),
		},
		InnerSeries: priceSeries,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			TickPosition: chart.TickPositionBetweenTicks,
			// ensure we render date timestamps with minute granularity
			ValueFormatter: chart.TimeMinuteValueFormatter,
		},
		Series: []chart.Series{
			bbSeries,
			priceSeries,
			hourlySMASeries,
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
