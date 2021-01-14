package dashboard

import (
	"context"
	"net/http"

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
