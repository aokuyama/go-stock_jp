package calendar_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/calendar"

	"github.com/stretchr/testify/assert"
)

func TestEnabledDate(t *testing.T) {
	var d *Date
	var err error
	d, err = NewDate("2020-10-01", true)
	assert.NoError(t, err)
	assert.Equal(t, `{"date":"2020-10-01","is_holiday":true}`, d.String(), "有効な日付")
	assert.True(t, d.IsHoliday())
	d, err = NewDate("2020-01-07", false)
	assert.NoError(t, err)
	assert.Equal(t, `{"date":"2020-01-07","is_holiday":false}`, d.String(), "有効な日付")
	assert.False(t, d.IsHoliday())
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
	assert.True(t, d1.IsEqualDate(d2), "休日設定は無視")
	assert.False(t, d1.IsEqual(d2), "休日設定はまで含める比較")
	d3, _ := NewDate("2020-01-01", true)
	assert.True(t, d1.IsEqualDate(d3))
	assert.True(t, d1.IsEqualDate(d3), "休日設定まであっていれば等価")
}
