package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledDate(t *testing.T) {
	var d *Date
	var err error
	d, err = NewDate("2020-10-01", true)
	assert.NoError(t, err)
	assert.Equal(t, "2020-10-01", d.date.String(), "有効な日付")
	assert.True(t, d.isHoliday)
	d, err = NewDate("2020-01-07", false)
	assert.NoError(t, err)
	assert.Equal(t, "2020-01-07", d.date.String(), "有効な日付")
	assert.False(t, d.isHoliday)
}

func TestDisabledQuantity(t *testing.T) {
	var err error
	_, err = NewDate("2020-1-1", true)
	assert.Error(t, err)
	_, err = NewDate("2020/10/01", true)
	assert.Error(t, err)
}

func TestEqual(t *testing.T) {
	d1, _ := NewDate("2020-01-01", true)
	d2, _ := NewDate("2020-01-01", false)
	assert.True(t, d1.isEqualDate(d2), "休日設定は無視")
	assert.False(t, d1.isEqual(d2), "休日設定はまで含める比較")
	d3, _ := NewDate("2020-01-01", true)
	assert.True(t, d1.isEqualDate(d3))
	assert.True(t, d1.isEqualDate(d3), "休日設定まであっていれば等価")
}
