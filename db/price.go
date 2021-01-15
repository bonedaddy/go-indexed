package db

import (
	"errors"
	"math"
	"time"
)

// Price is a given price entry for an asset
type Price struct {
	Model
	Type     Asset `gorm:"varchar(255)"`
	USDPrice float64
}

// RecordPrice records the given asset price in the database
func (d *Database) RecordPrice(asset string, price float64) error {
	if !IsValidAsset(asset) {
		return ErrInvalidAsset
	}
	return d.db.Create(&Price{Type: ToAsset(asset), USDPrice: price}).Error
}

// LastPrice returns the last recorded price
func (d *Database) LastPrice(asset string) (float64, error) {
	if !IsValidAsset(asset) {
		return 0, ErrInvalidAsset
	}
	var price Price
	if err := d.db.Model(&Price{}).Where("type = ?", asset).Last(&price).Error; err != nil {
		return 0, err
	}
	return price.USDPrice, nil
}

// GetAllPrices returns all price entries for a given asset
func (d *Database) GetAllPrices(asset string) ([]*Price, error) {
	if !IsValidAsset(asset) {
		return nil, ErrInvalidAsset
	}
	var prices []*Price
	return prices, d.db.Model(&Price{}).Where("type = ?", asset).Find(&prices).Error
}

// PricesInRange returns all prices items in the given range
func (d *Database) PricesInRange(asset string, windowInDays int) ([]*Price, error) {
	if !IsValidAsset(asset) {
		return nil,
			ErrInvalidAsset
	}
	return d.windowRangeQuery(asset, windowInDays)
}

// PriceAvgInRange returns the average price of the given asset during the last N days
func (d *Database) PriceAvgInRange(asset string, windowInDays int) (float64, error) {
	if !IsValidAsset(asset) {
		return 0, ErrInvalidAsset
	}
	prices, err := d.windowRangeQuery(asset, windowInDays)
	if err != nil {
		return 0, err
	}
	var totalPrice float64
	for _, price := range prices {
		totalPrice += price.USDPrice
	}
	return totalPrice / float64(len(prices)), nil
}

// PriceChangeInRange returns the price change percentage in the last N days
func (d *Database) PriceChangeInRange(asset string, windowInDays int) (float64, error) {
	if !IsValidAsset(asset) {
		return 0, ErrInvalidAsset
	}
	prices, err := d.windowRangeQuery(asset, windowInDays)
	if err != nil {
		return 0, err
	}
	switch len(prices) {
	case 0:
		return 0, errors.New("no prices found")
	case 1:
		return 0, nil // no price change
	default:
	}
	startPrice := prices[0].USDPrice
	finalPrice := prices[len(prices)-1].USDPrice
	// get the percentage change
	percentChange := ((finalPrice - startPrice) / math.Abs(startPrice))
	return percentChange, nil
}

func (d *Database) windowRangeQuery(asset string, windowInDays int) ([]*Price, error) {
	windowEnd := time.Now()
	windowStart := windowEnd.AddDate(0, 0, -windowInDays)
	var prices []*Price
	return prices, d.db.Model(&Price{}).Where(
		"type = ? AND created_at BETWEEN ? AND ?",
		asset, windowStart, windowEnd,
	).Find(&prices).Error
}
