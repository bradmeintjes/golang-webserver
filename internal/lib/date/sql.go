package date

import (
	"database/sql/driver"
	"errors"
	"time"
)

// Value implements of valuer for database/sql
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// Scan implements the database/sql scanner interface
func (d *Date) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		return errors.New("failed to scan null date")
	}

	if v, ok := value.(time.Time); ok {
		*d = Date(v.UTC())
		return nil
	} else {
		return errors.New("failed to scan date: cannot cast")
	}
}
