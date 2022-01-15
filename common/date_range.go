package common

import "errors"

type DateRange struct {
	begin Date
	end   Date
}

func NewDateRange(begin string, end string) (*DateRange, error) {
	b, err := NewDate(begin)
	if err != nil {
		return nil, err
	}
	e, err := NewDate(end)
	if err != nil {
		return nil, err
	}
	if b.Time().After(e.Time()) {
		return nil, errors.New("invalid date range:" + begin + " ~ " + end)
	}
	return &DateRange{
		begin: *b,
		end:   *e,
	}, nil
}

func (d *DateRange) String() string {
	return d.begin.String() + " ~ " + d.end.String()
}

func (d *DateRange) IsEqual(date *Date) bool {
	return d.String() == date.String()
}
