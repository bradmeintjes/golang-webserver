package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00001, down00001)
}

func up00001(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table todo (
			id				integer 	primary key autoincrement,
			content			text		not null,
			complete		integer		not null,
			created			integer		not null,
			completed_on	integer
		);
	`)
	return err
}

func down00001(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop table todo;
	`)
	return err
}
