package dashboard

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/db"
	"github.com/bonedaddy/go-indexed/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// CC10
	cc10TVL = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "pool",
		Name:      "cc10_tvl",
	})
	cc10TotalSupply = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "pool",
		Name:      "cc10_total_supply",
	})
	cc10DaiPrice = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "pool",
		Name:      "cc10_dai_price",
	})
	// DEFI5
	defi5TVL = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "pool",
		Name:      "defi5_tvl",
	})
	defi5TotalSupply = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "pool",
		Name:      "defi5_total_supply",
	})
	defi5DaiPrice = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "pool",
		Name:      "defi5_dai_price",
	})
	// NDX
	ndxTotalSupply = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "governance",
		Name:      "ndx_total_supply",
	})
	ndxDaiPrice = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "indexed",
		Subsystem: "governance",
		Name:      "ndx_dai_price",
	})
)

func init() {
	prometheus.MustRegister(cc10TVL, cc10TotalSupply, cc10DaiPrice, defi5TVL, defi5TotalSupply, defi5DaiPrice, ndxTotalSupply, ndxDaiPrice)
}

// ServePrometheusMetrics starts a http server to expose prometheus metrics
func ServePrometheusMetrics(ctx context.Context, listenAddr string) error {
	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Handler: handler,
		Addr:    listenAddr,
	}
	go func() {
		<-ctx.Done()
		srv.Close()
	}()
	return srv.ListenAndServe()
}

// UpdateMetrics is used to update the prometheus metrics
func UpdateMetrics(ctx context.Context, database *db.Database, bc *bclient.Client) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// set price information
			if price, err := database.LastPrice("ndx"); err == nil {
				ndxDaiPrice.Set(price)
			}
			if price, err := database.LastPrice("defi5"); err == nil {
				defi5DaiPrice.Set(price)
			}
			if price, err := database.LastPrice("cc10"); err == nil {
				cc10DaiPrice.Set(price)
			}
			// set tvl information
			if tvl, err := database.LastValueLocked("cc10"); err == nil {
				cc10TVL.Set(tvl)
			}
			if tvl, err := database.LastValueLocked("defi5"); err == nil {
				defi5TVL.Set(tvl)
			}
			// set total supply information
			if erc, err := bc.NDX(); err == nil {
				if supply, err := erc.TotalSupply(nil); err == nil {
					supplyF, _ := utils.ToDecimal(supply, 18).Float64()
					ndxTotalSupply.Set(supplyF)
				} else {
					log.Println("failed to get total supply: ", err)
				}
			}
			if erc, err := bc.DEFI5(); err == nil {
				if supply, err := erc.TotalSupply(nil); err == nil {
					supplyF, _ := utils.ToDecimal(supply, 18).Float64()
					defi5TotalSupply.Set(supplyF)
				} else {
					log.Println("failed to get total supply: ", err)
				}
			}
			if erc, err := bc.CC10(); err == nil {
				if supply, err := erc.TotalSupply(nil); err == nil {
					supplyF, _ := utils.ToDecimal(supply, 18).Float64()
					cc10TotalSupply.Set(supplyF)
				} else {
					log.Println("failed to get total supply: ", err)
				}
			}
		}
	}
}
