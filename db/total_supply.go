package db

// TotalSupply is used to track the total supply of a given asset
type TotalSupply struct {
	Model
	Type   string
	Supply float64
}

// RecordTotalSupply is used to record the supply for an asset
func (d *Database) RecordTotalSupply(asset string, supply float64) error {
	return d.db.Create(&TotalSupply{Type: asset, Supply: supply}).Error
}

// LastTotalSupply returns the last recorded total supply for an asset
func (d *Database) LastTotalSupply(asset string) (float64, error) {
	var ts TotalSupply
	if err := d.db.Model(&TotalSupply{}).Where(
		"type = ?", asset).Last(&ts).Error; err != nil {
		return 0, err
	}
	return ts.Supply, nil
}
