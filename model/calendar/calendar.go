package calendar

import (
	"errors"
	"sort"

	"github.com/aokuyama/go-stock_jp/common"
)

type Calendar struct {
	dates []*Date
}

func NewCalendar(dates *[]*Date) (*Calendar, error) {
	var err error
	c := Calendar{}
	err = c.setDates(dates)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type ByDate []*Date

func (d ByDate) Len() int {
	return len(d)
}
func (d ByDate) Less(i, j int) bool {
	return d[i].date.String() < d[j].date.String()
}
func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (c *Calendar) setDates(dates *[]*Date) error {
	if dates == nil {
		return nil
	}
	for _, d := range *dates {
		for _, sd := range c.dates {
			if d.isEqualDate(sd) {
				return errors.New("duplicate date:" + d.date.String())
			}
		}
		c.dates = append(c.dates, d)
	}
	sort.Sort(ByDate(c.dates))
	return nil
}

func (c *Calendar) TradeDays() *[]*common.Date {
	var dates []*common.Date
	for _, d := range c.dates {
		if d.IsTradeDay() {
			dates = append(dates, &d.date)
		}
	}
	return &dates
}

func (c *Calendar) TradeDayOnAfterTomorrow(today *common.Date) *common.Date {
	for _, d := range *c.TradeDays() {
		if d.After(today) {
			return d
		}
	}
	return nil
}

func (c *Calendar) TradeDayOnAfterToday(today *common.Date) *common.Date {
	for _, d := range *c.TradeDays() {
		if !d.Before(today) {
			return d
		}
	}
	return nil
}
