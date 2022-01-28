package common

import (
	"time"
)

type Date time.Time

func NewDate(v string) (*Date, error) {
	t, err := time.Parse(time.RFC3339, v+"T00:00:00+00:00")
	if err != nil {
		return nil, err
	}
	d := Date(t)
	return &d, nil
}

func NewDateAdded(date *Date, plus time.Duration) *Date {
	t := date.Time()
	d := Date(t.Add(plus))
	return &d
}

func (d *Date) String() string {
	return time.Time(*d).Format("2006-01-02")
}

func (d *Date) Time() time.Time {
	return time.Time(*d)
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Date) IsEqual(date *Date) bool {
	return d.String() == date.String()
}

func (d *Date) Before(date *Date) bool {
	return d.Time().Before(date.Time())
}

func (d *Date) After(date *Date) bool {
	return d.Time().After(date.Time())
}
