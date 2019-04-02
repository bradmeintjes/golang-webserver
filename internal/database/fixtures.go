package database

import (
	"go/webserver/internal/util/osutil"
	"go/webserver/internal/util/testutil"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Fixture returns a temporary test database for testing.
func Fixture(file string) (db *sqlx.DB, clean func(), err error) {
	const driver = "sqlite3"
	var cleanup testutil.Cleanup

	defer cleanup.Recover()
	cleanup.Add(func() { os.Remove(file) })

	clean = cleanup.Run
	err = osutil.EnsureFilePresent(file, os.ModePerm)
	if err != nil {
		return
	}

	db, err = New(Config{
		Driver:  driver,
		ConnStr: file,
	})
	return
}
