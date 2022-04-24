package common_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/common"
	"github.com/stretchr/testify/assert"
)

func TestEnabledDateRange(t *testing.T) {
	var d *DateRange
	var err error
	d, err = NewDateRange("2000-01-01", "2000-01-02")
	assert.Equal(t, "2000-01-01 ~ 2000-01-02", d.String(), "有効な日付範囲")
	assert.Equal(t, "2000-01-01", d.Begin().String())
	assert.Equal(t, "2000-01-02", d.End().String())
	assert.NoError(t, err)
	d, err = NewDateRange("2000-01-01", "2020-01-02")
	assert.Equal(t, "2000-01-01 ~ 2020-01-02", d.String(), "有効な日付範囲")
	assert.NoError(t, err)
	d, err = NewDateRange("2020-01-02", "2020-01-02")
	assert.Equal(t, "2020-01-02 ~ 2020-01-02", d.String(), "同じ日付でもOK")
	assert.NoError(t, err)
}

func TestDisabledDateRange(t *testing.T) {
	var d *DateRange
	var err error
	d, err = NewDateRange("1", "2000-01-01")
	assert.Error(t, err, "どちらかの日付が無効")
	assert.Nil(t, d)
	d, err = NewDateRange("2000-01-01", "1")
	assert.Error(t, err, "どちらかの日付が無効")
	assert.Nil(t, d)
	d, err = NewDateRange("2000-01-02", "2000-01-01")
	assert.Error(t, err, "開始日付の方が後になっている")
	assert.Nil(t, d)
}

func TestDateRangeToList(t *testing.T) {
	var d *DateRange
	var err error
	d, err = NewDateRange("2000-01-01", "2000-01-08")
	assert.NoError(t, err)
	assert.Equal(t, 8, len(d.ToList()))
	dates := d.ToList()
	for i, date := range dates {
		if i > 0 {
			assert.True(t, date.After(dates[i-1]))
			assert.False(t, date.IsEqual(dates[i-1]))
		}
	}
	assert.Equal(t, "2000-01-01", dates[0].String())
	assert.Equal(t, "2000-01-08", dates[7].String())
	d, err = NewDateRange("2010-09-28", "2010-10-02")
	assert.NoError(t, err)
	assert.Equal(t, 5, len(d.ToList()))
	d, err = NewDateRange("2020-01-02", "2020-01-02")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(d.ToList()))
}
