package db

import "testing"

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
}
