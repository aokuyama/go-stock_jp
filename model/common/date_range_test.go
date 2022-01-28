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
