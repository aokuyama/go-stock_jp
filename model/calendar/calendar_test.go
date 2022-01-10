package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateSorted(t *testing.T) {
	d1, _ := newDate("2021-04-01", true)
	d2, _ := newDate("2020-10-10", false)
	d3, _ := newDate("2020-10-08", true)
	d4, _ := newDate("2020-10-09", false)
	d5, _ := newDate("2019-12-15", true)
	var dates []*Date
	dates = append(dates, d1, d2, d3, d4, d5)
	c, err := NewCalendar(&dates)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(c.dates))
	assert.Equal(t, "2019-12-15", c.dates[0].date.String())
	assert.Equal(t, "2020-10-08", c.dates[1].date.String())
	assert.Equal(t, "2020-10-09", c.dates[2].date.String())
	assert.Equal(t, "2020-10-10", c.dates[3].date.String())
	assert.Equal(t, "2021-04-01", c.dates[4].date.String())
}

func TestErrorDuplicateDates(t *testing.T) {
	d1, _ := newDate("2021-05-05", true)
	d2, _ := newDate("2021-05-05", false)
	var dates []*Date
	dates = append(dates, d1, d2)
	_, err := NewCalendar(&dates)
	assert.Error(t, err)
}
