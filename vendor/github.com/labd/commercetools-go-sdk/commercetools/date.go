package commercetools

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date holds date information for Commercetools API format
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// NewDate initializes a Date struct
func NewDate(year int, month time.Month, day int) Date {
	return Date{Year: year, Month: month, Day: day}
}

// MarshalJSON marshals into the commercetools date format
func (d *Date) MarshalJSON() ([]byte, error) {
	value := fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
	return []byte(value), nil
}

// UnmarshalJSON decodes JSON data into a Date struct
func (d *Date) UnmarshalJSON(data []byte) error {
	var input string
	err := json.Unmarshal(data, &input)
	if err != nil {
		return err
	}

	value, err := time.Parse("2006-01-02", input)
	if err != nil {
		return err
	}

	d.Year = value.Year()
	d.Month = value.Month()
	d.Day = value.Day()
	return nil
}
