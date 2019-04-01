package date

import "database/sql/driver"

type NullDate struct {
	Date  Date
	Valid bool
}

// Scan implements the Scanner interface.
func (nd *NullDate) Scan(value interface{}) error {
	if value == nil {
		nd.Date, nd.Valid = New(0, 0, 0), false
		return nil
	}
	nd.Valid = true

	return nd.Date.Scan(value)
}

// Value implements the driver Valuer interface.
func (nd NullDate) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Date.String(), nil
}
