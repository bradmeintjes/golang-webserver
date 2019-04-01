package database

import (
	"github.com/jmoiron/sqlx"
)

type TxFn func(tx *sqlx.Tx) error

// AsTransaction will execute the given TxFn using a single db transaction
// and will rollback or commit based on the error returned from the fn
func AsTransaction(db *sqlx.DB, fn TxFn) (err error) {
	var tx *sqlx.Tx
	tx, err = db.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
