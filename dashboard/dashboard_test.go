package dashboard

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestServePrometheusMetrics(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := ServePrometheusMetrics(ctx, ":12345")
	if err != nil && err != http.ErrServerClosed {
		t.Fatal(err)
	}
}
