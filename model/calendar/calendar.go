package calendar

import (
	"errors"
	"sort"
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
