package db

import (
	"time"

	"gorm.io/gorm"
)

// Model is a copy of the gorm model type but using slightly different configurations
type Model struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"` // this is changed to include the index
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
