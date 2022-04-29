package calendar_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/calendar"
	"github.com/stretchr/testify/assert"
)

func TestEnabledTradeDayMode(t *testing.T) {
	var m *TradeDayMode
	var err error
	m, err = NewTradeDayMode("today")
	assert.NoError(t, err)
	assert.Equal(t, "today", m.String(), "取引日モード：当日以降の取引日")
	assert.True(t, m.IsTradeDayOnAfterToday())
	assert.False(t, m.IsTradeDayOnAfterTomorrow())
	m, err = NewTradeDayMode("tomorrow")
	assert.NoError(t, err)
	assert.Equal(t, "tomorrow", m.String(), "取引日モード：翌日以降の取引日")
	assert.False(t, m.IsTradeDayOnAfterToday())
	assert.True(t, m.IsTradeDayOnAfterTomorrow())
}

func TestDisabledTradeDayMode(t *testing.T) {
	var m *TradeDayMode
	var err error
	m, err = NewTradeDayMode("2000-01-01")
	assert.Nil(t, m)
	assert.Error(t, err)
	m, err = NewTradeDayMode("1")
	assert.Nil(t, m)
	assert.Error(t, err)
}
