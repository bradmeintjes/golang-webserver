package date

import (
	"time"
)

// Date is a simple wrapper around a time which removes the time portion and keeps only
// the local date portion
type Date time.Time

const dateFmt = `"2006-01-02"`

// New creates a new date instance using the given year, month and day
// month is 1 based as it is within the time package
func New(year int, month time.Month, day int) Date {
	return Date(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	dt, err := Parse(string(b))

	if err == nil {
		*d = *dt
	}
	return err
}

func (d Date) String() string {
	return time.Time(d).Format(dateFmt)
}

func Parse(s string) (*Date, error) {
	t, err := time.Parse(dateFmt, s)
	if err != nil {
		return nil, err
	}

	d := Date(t.UTC())
	return &d, nil
}
