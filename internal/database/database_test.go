package database

import (
	"go/webserver/internal/util/testutil"
	"testing"

	_ "github.com/lib/pq"
)

func TestMigrations(t *testing.T) {
	db, f, err := Fixture("test.db")
	if f != nil {
		defer f()
	}

	testutil.AssertNoErr(t, err, "database should migrate")
	testutil.AssertNotNil(t, db, "")
}
