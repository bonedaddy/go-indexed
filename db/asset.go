package db

import (
	"errors"
	"strings"
)

// Asset is a given crypto asset tracked by it's Ticker symbol
type Asset string

// String returns the Asset value in string type
func (a Asset) String() string {
	return string(a)
}

var (
	// Assets are all currently supported assets for which we track prices
	Assets = []Asset{
		Asset("NDX"),
		Asset("DEFI5"),
		Asset("CC10"),
		Asset("ETH"),
		Asset("ORCL5"),
		Asset("DEGEN10"),
	}
	// ErrInvalidAsset is an error returned when the given asset specified is invalid
	ErrInvalidAsset = errors.New("invalid asset")
)

// IsValidAsset determines whether or not the given asset is one we are tracking
func IsValidAsset(asset string) bool {
	switch strings.ToLower(asset) {
	case "ndx", "defi5", "cc10", "orcl5", "eth":
		return true
	default:
		return false
	}
}

// ToAsset is a helper function to type convert string -> Asset
func ToAsset(asset string) Asset {
	return Asset(asset)
}
