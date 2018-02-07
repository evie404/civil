package civil

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Scan implements the Scanner interface.
// The value type must be time.Time or string / []byte (formatted time-string),
// otherwise Scan fails.
func (date *Date) Scan(value interface{}) (err error) {
	if value == nil {
		date.Year = 0
		date.Month = 0
		date.Day = 0
		return
	}

	switch v := value.(type) {
	case time.Time:
		date.Year, date.Month, date.Day = v.Date()
		return
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		date.Year, date.Month, date.Day = t.Date()
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		date.Year, date.Month, date.Day = t.Date()
		return nil
	}

	return fmt.Errorf("Can't convert %T to civil.Date", value)
}

// Value implements the driver Valuer interface.
func (date Date) Value() (driver.Value, error) {
	if !date.IsValid() {
		return nil, nil
	}
	return date.String(), nil
}
