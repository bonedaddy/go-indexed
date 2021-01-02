package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrice(t *testing.T) {
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
}
