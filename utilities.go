package gotick

import (
	"encoding/json"
	"time"
)

type TickDate struct{ time.Time }

func (d *TickDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
