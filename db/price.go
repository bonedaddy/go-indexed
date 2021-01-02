package db

import (
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
	err := d.db.Model(&Price{}).Where("type = ?", asset).Last(&price).Error
	if err != nil {
		return 0, err
	}
	return price.USDPrice, nil
}
