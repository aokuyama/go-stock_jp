package calendar

import (
	"encoding/json"

	"github.com/aokuyama/go-stock_jp/model/common"
)

type Date struct {
	date      common.Date
	isHoliday bool
}

type dateJson struct {
	Date      common.Date `json:"date"`
	IsHoliday bool        `json:"is_holiday"`
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

func (d *Date) IsEqual(date *Date) bool {
	return (d.isHoliday == date.isHoliday) && d.IsEqualDate(date)
}

func (d *Date) IsEqualDate(date *Date) bool {
	return d.date.IsEqual(&date.date)
}

func (d *Date) IsHoliday() bool {
	return d.isHoliday
}

func (d *Date) IsTradeDay() bool {
	return !d.IsHoliday()
}

func (d *Date) String() string {
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(&dateJson{
		Date:      d.date,
		IsHoliday: d.isHoliday,
	})
}

func (d *Date) UnmarshalJSON(b []byte) error {
	j := dateJson{}
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*d = Date{
		date:      j.Date,
		isHoliday: j.IsHoliday,
	}
	return nil
}
