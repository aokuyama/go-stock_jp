package calendar

import "github.com/aokuyama/go-stock_jp/common"

type Date struct {
	date      common.Date
	isHoliday bool
}

func NewDate(date string, isHoliday bool) (*Date, error) {
	d, err := common.NewDate(date)
	if err != nil {
		return nil, err
	}
	return &Date{
		date:      *d,
		isHoliday: isHoliday,
	}, nil
}

func (d *Date) isEqual(date *Date) bool {
	return (d.isHoliday == date.isHoliday) && d.isEqualDate(date)
}

func (d *Date) isEqualDate(date *Date) bool {
	return d.date.IsEqual(&date.date)
}

func (d *Date) IsTradeDay() bool {
	return !d.isHoliday
}
