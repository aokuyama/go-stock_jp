package calendar_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/calendar"

	"github.com/aokuyama/go-stock_jp/common"
	"github.com/stretchr/testify/assert"
)

func TestDateSorted(t *testing.T) {
	d1, _ := NewDate("2021-04-01", true)
	d2, _ := NewDate("2020-10-10", false)
	d3, _ := NewDate("2020-10-08", true)
	d4, _ := NewDate("2020-10-09", false)
	d5, _ := NewDate("2019-12-15", true)
	var dates []*Date
	dates = append(dates, d1, d2, d3, d4, d5)
	c, err := New(&dates)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(c.Dates))
	assert.Equal(t, "2019-12-15", c.Dates[0].Date.String())
	assert.Equal(t, "2020-10-08", c.Dates[1].Date.String())
	assert.Equal(t, "2020-10-09", c.Dates[2].Date.String())
	assert.Equal(t, "2020-10-10", c.Dates[3].Date.String())
	assert.Equal(t, "2021-04-01", c.Dates[4].Date.String())
}

func TestErrorDuplicateDates(t *testing.T) {
	d1, _ := NewDate("2021-05-05", true)
	d2, _ := NewDate("2021-05-05", false)
	var dates []*Date
	dates = append(dates, d1, d2)
	_, err := New(&dates)
	assert.Error(t, err)
}

func TestPickupTradeDays(t *testing.T) {
	d1, _ := NewDate("2021-05-05", true)
	d2, _ := NewDate("2021-05-06", false)
	d3, _ := NewDate("2021-05-07", true)
	d4, _ := NewDate("2021-05-08", false)
	d5, _ := NewDate("2021-05-09", false)
	d6, _ := NewDate("2021-05-10", true)
	var dates []*Date
	dates = append(dates, d1, d2, d3, d4, d5, d6)
	c, _ := New(&dates)
	ts := c.TradeDays()
	assert.Equal(t, 3, len(*ts))
	assert.Equal(t, "2021-05-06", (*ts)[0].String())
	assert.Equal(t, "2021-05-08", (*ts)[1].String())
	assert.Equal(t, "2021-05-09", (*ts)[2].String())
}
func TestTradeDayOnAfterTomorrow(t *testing.T) {
	var d *common.Date
	c := newCalendar()
	d, _ = common.NewDate("2022-01-10")
	assert.Equal(t, "2022-01-11", c.TradeDayOnAfterTomorrow(d).String())
	d, _ = common.NewDate("2022-01-11")
	assert.Equal(t, "2022-01-12", c.TradeDayOnAfterTomorrow(d).String())
	d, _ = common.NewDate("2022-01-14")
	assert.Equal(t, "2022-01-17", c.TradeDayOnAfterTomorrow(d).String())
	d, _ = common.NewDate("2022-01-15")
	assert.Equal(t, "2022-01-17", c.TradeDayOnAfterTomorrow(d).String())
}

func newCalendar() *Calendar {
	d1, _ := NewDate("2022-01-11", false)
	d2, _ := NewDate("2022-01-12", false)
	d3, _ := NewDate("2022-01-13", false)
	d4, _ := NewDate("2022-01-14", false)
	d5, _ := NewDate("2022-01-15", true)
	d6, _ := NewDate("2022-01-16", true)
	d7, _ := NewDate("2022-01-17", false)
	var dates []*Date
	dates = append(dates, d1, d2, d3, d4, d5, d6, d7)
	c, _ := New(&dates)
	return c
}

func TestTradeDayOnAfterToday(t *testing.T) {
	var d *common.Date
	c := newCalendar()
	d, _ = common.NewDate("2022-01-10")
	assert.Equal(t, "2022-01-11", c.TradeDayOnAfterToday(d).String())
	d, _ = common.NewDate("2022-01-11")
	assert.Equal(t, "2022-01-11", c.TradeDayOnAfterToday(d).String())
	d, _ = common.NewDate("2022-01-14")
	assert.Equal(t, "2022-01-14", c.TradeDayOnAfterToday(d).String())
	d, _ = common.NewDate("2022-01-15")
	assert.Equal(t, "2022-01-17", c.TradeDayOnAfterToday(d).String())
}
