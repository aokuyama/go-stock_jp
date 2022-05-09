package calendar

import (
	"encoding/json"
	"errors"
	"sort"

	"github.com/aokuyama/go-stock_jp/model/common"
)

type Calendar struct {
	dates []*Date
}

type calendarJson struct {
	Dates []*Date `json:"dates"`
}

func New(dates *[]*Date) (*Calendar, error) {
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
			if d.IsEqualDate(sd) {
				return errors.New("duplicate date:" + d.date.String())
			}
		}
		c.dates = append(c.dates, d)
	}
	sort.Sort(ByDate(c.dates))
	return nil
}

func (c *Calendar) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func (c *Calendar) MarshalJSON() ([]byte, error) {
	return json.Marshal(&calendarJson{
		Dates: c.dates,
	})
}

func (c *Calendar) UnmarshalJSON(b []byte) error {
	j := calendarJson{}
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*c = Calendar{
		dates: j.Dates,
	}
	return nil
}

func (c *Calendar) Len() int {
	return len(c.dates)
}

func (c *Calendar) Nth(i int) Date {
	return *c.dates[i]
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
