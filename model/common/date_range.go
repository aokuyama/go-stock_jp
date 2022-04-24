package common

import (
	"errors"
	"time"
)

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
	return NewDateRangeByDates(b, e)
}

func NewDateRangeByDates(begin *Date, end *Date) (*DateRange, error) {
	if begin.Time().After(end.Time()) {
		return nil, errors.New("invalid date range:" + begin.String() + " ~ " + end.String())
	}
	return &DateRange{
		begin: *begin,
		end:   *end,
	}, nil
}

func (d *DateRange) Begin() *Date {
	return &d.begin
}

func (d *DateRange) End() *Date {
	return &d.end
}

func (d *DateRange) String() string {
	return d.begin.String() + " ~ " + d.end.String()
}

func (d *DateRange) IsEqual(date *Date) bool {
	return d.String() == date.String()
}

func (d *DateRange) ToList() []*Date {
	var dates []*Date
	var date *Date
	date, _ = NewDate(d.begin.String())
	for {
		dates = append(dates, date)
		if date.IsEqual(&d.end) {
			return dates
		}
		date = NewDateAdded(date, time.Hour*24)
	}
}
