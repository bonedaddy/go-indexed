package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotalValueLocked(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("indexed.db")
	})
	db := newTestDB(t)

	type args struct {
		pool      string
		price     float64
		wantPrice float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"DEFI5-5", args{"defi5", 5, 5}},
		{"DEFI5-10", args{"defi5", 10, 10}},
		{"CC10-15", args{"cc10", 15, 15}},
		{"CC10-20", args{"cc10", 20, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, db.RecordValueLocked(tt.args.pool, tt.args.price))
			price, err := db.LastValueLocked(tt.args.pool)
			require.NoError(t, err)
			require.Equal(t, price, float64(tt.args.wantPrice))
		})
	}
}
