package db

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

// Opts is used to configure database connections
type Opts struct {
	Type           string
	Host           string
	Port           string
	User           string
	Password       string
	DBName         string
	DBPath         string
	SSLModeDisable bool
}

// DSN returns the config string used with gorm
func (db *Opts) DSN() string {
	var sslModeDisable string
	if db.SSLModeDisable {
		sslModeDisable = "sslmode=disable"
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s %s",
		db.Host, db.User, db.Password, db.DBName, db.Port, sslModeDisable,
	)
}

// Open returns the gorm dialector for the corresponding db type
func (db *Opts) Open() (gorm.Dialector, error) {
	switch strings.ToLower(db.Type) {
	case "postgres":
		return postgres.Open(db.DSN()), nil
	case "sqlite":
		return sqlite.Open(db.DBPath + db.DBName + ".db"), nil
	default:
		return nil, errors.New("unsupported db type")
	}

}

// New returns a new database client
func New(opts *Opts) (*Database, error) {
	dialector, err := opts.Open()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

// AutoMigrate is used to automatically migrate datbase tables
func (d *Database) AutoMigrate() error {
	var tables []interface{}
	tables = append(tables, &Price{}, &TotalValueLocked{})
	for _, table := range tables {
		if err := d.db.AutoMigrate(table); err != nil {
			return err
		}
	}
	return nil
}

// Close shuts down the database
func (d *Database) Close() error {
	sdb, err := d.db.DB()
	if err != nil {
		return err
	}
	return sdb.Close()
}
