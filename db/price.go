package db

import (
	"time"

	"gorm.io/gorm"
)

// Price is a given price entry for an asset
type Price struct {
	gorm.Model
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

// PriceAvgInRange returns the average price of the given asset during the last N days
func (d *Database) PriceAvgInRange(asset string, windowInDays int) (float64, error) {
	if !IsValidAsset(asset) {
		return 0, ErrInvalidAsset
	}
	windowEnd := time.Now()
	windowStart := windowEnd.AddDate(0, 0, -windowInDays)
	var prices []*Price
	if err := d.db.Model(&Price{}).Where(
		"type = ? AND created_at BETWEEN ? AND ?",
		asset, windowStart, windowEnd,
	).Find(&prices).Error; err != nil {
		return 0, err
	}
	var totalPrice float64
	for _, price := range prices {
		totalPrice += price.USDPrice
	}
	return totalPrice / float64(len(prices)), nil
}
