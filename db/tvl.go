package db

// TotalValueLocked records the USD value of funds locked within index pool
type TotalValueLocked struct {
	Model
	PoolName    string
	ValueLocked float64
}

// RecordValueLocked records the USD value locked into a pool
func (d *Database) RecordValueLocked(pool string, locked float64) error {
	return d.db.Create(&TotalValueLocked{PoolName: pool, ValueLocked: locked}).Error
}

// LastValueLocked returns the latest record total value locked entry
func (d *Database) LastValueLocked(pool string) (float64, error) {
	var tvl TotalValueLocked
	if err := d.db.Model(&TotalValueLocked{}).Where(
		"pool_name = ?", pool).Last(&tvl).Error; err != nil {
		return 0, err
	}
	return tvl.ValueLocked, nil
}
