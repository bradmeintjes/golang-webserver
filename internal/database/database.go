package database

import (
	"fmt"

	_ "go/webserver/internal/database/migrations" // Add migrations.

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

// New creates and initialises a new database
func New(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(config.Driver, config.ConnStr)
	if err != nil {
		return nil, fmt.Errorf("open database: %s", err)
	}

	if err := goose.SetDialect(config.Driver); err != nil {
		return nil, fmt.Errorf("set dialect as %s: %s", config.Driver, err)
	}
	if err := goose.Up(db.DB, "."); err != nil {
		return nil, fmt.Errorf("perform db migration: %s", err)
	}
	return db, nil
}
