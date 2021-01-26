package discord

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseValue(t *testing.T) {
	type args struct {
		num     float64
		wantStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{"125-0", args{125, "125"}},
		{"125-1", args{125.12345, "125"}},
		{"500K-0", args{500000, "500K"}},
		{"500K-1", args{500000.12345, "500K"}},
		{"1M-0", args{1000000, "1M"}},
		{"1M-1", args{1000000.12345, "1M"}},
		{"10M-0", args{10000000, "10M"}},
		{"10M-1", args{10000000.12345, "10M"}},
		{"100M-0", args{100000000, "100M"}},
		{"100M-1", args{100000000.12345, "100M"}},
		{"1B-0", args{1000000000, "1B"}},
		{"1B-1", args{1000000000.12345, "1B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.args.wantStr, ParseValue(tt.args.num))
		})
	}
}
