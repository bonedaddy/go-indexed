package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrice(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("indexed.db")
	})
	db := newTestDB(t)
	t.Run("RecordPrice", func(t *testing.T) {
		type args struct {
			asset string
			price float64
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"NDX", args{"ndx", 1}, false},
			{"CC10", args{"cc10", 1}, false},
			{"DEFI5", args{"defi5", 1}, false},
			{"InvalidAsset", args{"invalidasset", 1}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := db.RecordPrice(tt.args.asset, tt.args.price)
				if (err != nil) != tt.wantErr {
					t.Fatalf("RecordPrice() err %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
	t.Run("LastPrice", func(t *testing.T) {
		type args struct {
			asset       string
			firstPrice  float64
			secondPrice float64
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"NDX", args{"ndx", 10, 11}, false},
			{"CC10", args{"cc10", 12, 13}, false},
			{"DEFI5", args{"defi5", 14, 15}, false},
			{"InvalidAsset", args{"invalidasset", 11, 11}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := db.RecordPrice(tt.args.asset, tt.args.firstPrice)
				if (err != nil) != tt.wantErr {
					t.Fatalf("RecordPrice() err %v, wantErr %v", err, tt.wantErr)
				}
				if tt.wantErr {
					return
				}
				price, err := db.LastPrice(tt.args.asset)
				require.NoError(t, err)
				require.Equal(t, price, tt.args.firstPrice)

				require.NoError(t, db.RecordPrice(tt.args.asset, tt.args.secondPrice))

				price, err = db.LastPrice(tt.args.asset)
				require.NoError(t, err)
				require.Equal(t, price, tt.args.secondPrice)
			})
		}
	})
	t.Run("GetAllPrices", func(t *testing.T) {
		type args struct {
			asset       string
			wantEntries int
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"NDX", args{"ndx", 3}, false},
			{"CC10", args{"cc10", 3}, false},
			{"DEFI5", args{"defi5", 3}, false},
			{"InvalidAsset", args{"invalidasset", 1}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				entries, err := db.GetAllPrices(tt.args.asset)
				if (err != nil) != tt.wantErr {
					t.Fatalf("GetAllPrices() err %v, wantErr %v", err, tt.wantErr)
				}
				if tt.wantErr {
					return
				}
				require.Len(t, entries, tt.args.wantEntries)
				priceAvg, err := db.PriceAvgInRange(tt.args.asset, 1)
				require.NoError(t, err)
				t.Log("price avg: ", priceAvg)
			})
		}
	})
	t.Run("GetAllPrices", func(t *testing.T) {
		type args struct {
			asset       string
			wantEntries int
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"NDX", args{"ndx", 3}, false},
			{"CC10", args{"cc10", 3}, false},
			{"DEFI5", args{"defi5", 3}, false},
			{"InvalidAsset", args{"invalidasset", 1}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				entries, err := db.GetAllPrices(tt.args.asset)
				if (err != nil) != tt.wantErr {
					t.Fatalf("GetAllPrices() err %v, wantErr %v", err, tt.wantErr)
				}
				if tt.wantErr {
					return
				}
				require.Len(t, entries, tt.args.wantEntries)
			})
		}
	})
	t.Run("PriceAvgInRange", func(t *testing.T) {
		type args struct {
			asset     string
			window    int
			wantPrice float64
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"NDX", args{"ndx", 1, 7.333333333333333}, false},
			{"CC10", args{"cc10", 1, 8.666666666666666}, false},
			{"DEFI5", args{"defi5", 1, 10}, false},
			{"InvalidAsset", args{"invalidasset", 1, 0}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				priceAvg, err := db.PriceAvgInRange(tt.args.asset, tt.args.window)
				if (err != nil) != tt.wantErr {
					t.Fatalf("GetAllPrices() err %v, wantErr %v", err, tt.wantErr)
				}
				if tt.wantErr {
					return
				}
				entries, err := db.GetAllPrices(tt.args.asset)
				require.NoError(t, err)
				var totalPrice float64
				for _, entry := range entries {
					totalPrice += entry.USDPrice
				}
				avgPrice := totalPrice / float64(len(entries))
				require.Equal(t, avgPrice, priceAvg)
				// ensure recording a new price changes the average
				require.NoError(t, db.RecordPrice(tt.args.asset, 19))
				newPriceAvg, err := db.PriceAvgInRange(tt.args.asset, tt.args.window)
				require.NotEqual(t, newPriceAvg, priceAvg)
			})
		}
	})
}
