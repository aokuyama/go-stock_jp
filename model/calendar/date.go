package calendar

import "github.com/aokuyama/go-stock_jp/model/common"

type Date struct {
	Date      common.Date
	IsHoliday bool
}

func NewDate(date string, isHoliday bool) (*Date, error) {
	d, err := common.NewDate(date)
	if err != nil {
		return nil, err
	}
	return &Date{
		Date:      *d,
		IsHoliday: isHoliday,
	}, nil
}

func (d *Date) IsEqual(date *Date) bool {
	return (d.IsHoliday == date.IsHoliday) && d.IsEqualDate(date)
}

func (d *Date) IsEqualDate(date *Date) bool {
	return d.Date.IsEqual(&date.Date)
}

func (d *Date) IsTradeDay() bool {
	return !d.IsHoliday
}
