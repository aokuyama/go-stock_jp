package service_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/service"

	"github.com/aokuyama/go-stock_jp/model/calendar"
	"github.com/aokuyama/go-stock_jp/model/calendar/mock"
	"github.com/aokuyama/go-stock_jp/model/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func getNewTestGetTradeDay(t *testing.T) *GetTradeDay {
	gomock.NewController(t)
	mc := gomock.NewController(t)
	defer mc.Finish()
	repo := mock.NewMockCalendarRepository(mc)
	s, err := NewGetTradeDay(repo)
	if err != nil {
		panic(err)
	}
	return s
}

func TestGetTradeDay(t *testing.T) {
	s := getNewTestGetTradeDay(t)
	ex, _ := common.NewDateRange("2022-01-16", "2022-01-30")
	s.Repository.(*mock.MockCalendarRepository).EXPECT().
		LoadByDateRange(ex).
		Return(newCalendar(), nil)
	t.Run("after tomorrow", func(t *testing.T) {
		r, _ := s.GetTradeDayOnAfterTomorrow("2022-01-16")
		assert.Equal(t, "2022-01-18", r.String())
	})
	s.Repository.(*mock.MockCalendarRepository).EXPECT().
		LoadByDateRange(ex).
		Return(newCalendar(), nil)
	t.Run("after today", func(t *testing.T) {
		r, _ := s.GetTradeDayOnAfterToday("2022-01-16")
		assert.Equal(t, "2022-01-16", r.String())
	})
}

func newCalendar() *calendar.Calendar {
	d1, _ := calendar.NewDate("2022-01-16", false)
	d2, _ := calendar.NewDate("2022-01-17", true)
	d3, _ := calendar.NewDate("2022-01-18", false)
	var dates []*calendar.Date
	dates = append(dates, d1, d2, d3)
	c, _ := calendar.New(&dates)
	return c
}

func TestErrorGetTradeDayOnAfterTomorrow(t *testing.T) {
	s := getNewTestGetTradeDay(t)
	r, err := s.GetTradeDayOnAfterTomorrow("あ")
	assert.Error(t, err)
	assert.Nil(t, r)
	r2, err2 := s.GetTradeDayOnAfterToday("")
	assert.Error(t, err2)
	assert.Nil(t, r2)
	t.Run("not found", func(t *testing.T) {
		ex, _ := common.NewDateRange("2022-01-19", "2022-02-02")
		s.Repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex).
			Return(newCalendar(), nil)
		r, err := s.GetTradeDayOnAfterTomorrow("2022-01-19")
		assert.Error(t, err)
		assert.Nil(t, r)
	})
	t.Run("not found", func(t *testing.T) {
		ex, _ := common.NewDateRange("2022-01-19", "2022-02-02")
		s.Repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex).
			Return(newCalendar(), nil)
		r, err := s.GetTradeDayOnAfterToday("2022-01-19")
		assert.Error(t, err)
		assert.Nil(t, r)
	})
}

func TestGetDateByTradeDayMode(t *testing.T) {
	mockCal := func() *calendar.Calendar {
		d1, _ := calendar.NewDate("2022-04-27", false)
		d2, _ := calendar.NewDate("2022-04-28", false)
		d3, _ := calendar.NewDate("2022-04-29", true)
		d4, _ := calendar.NewDate("2022-04-30", true)
		d5, _ := calendar.NewDate("2022-05-01", true)
		d6, _ := calendar.NewDate("2022-05-02", false)
		d7, _ := calendar.NewDate("2022-05-03", true)
		var dates []*calendar.Date
		dates = append(dates, d1, d2, d3, d4, d5, d6, d7)
		c, _ := calendar.New(&dates)
		return c
	}
	var d *common.Date
	var err error
	s := getNewTestGetTradeDay(t)

	t.Run("static date", func(t *testing.T) {
		d, err = s.GetDateByTradeDayMode("2000-01-02")
		assert.NoError(t, err)
		assert.Equal(t, "2000-01-02", d.String())
		d, err = s.GetDateByTradeDayMode("2022-04-29")
		assert.NoError(t, err)
		assert.Equal(t, "2022-04-29", d.String())
	})

	t.Run("trade day on after today", func(t *testing.T) {
		ex, _ := common.NewDateRange("2022-04-28", "2022-05-12")
		s.Repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex).
			Return(mockCal(), nil)
		s.Today, _ = common.NewDate("2022-04-28")
		d, err = s.GetDateByTradeDayMode("today")
		assert.NoError(t, err)
		assert.Equal(t, "2022-04-28", d.String())

		ex2, _ := common.NewDateRange("2022-04-29", "2022-05-13")
		s.Repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex2).
			Return(mockCal(), nil)
		s.Today, _ = common.NewDate("2022-04-29")
		d, err = s.GetDateByTradeDayMode("today")
		assert.NoError(t, err)
		assert.Equal(t, "2022-05-02", d.String())
	})

	t.Run("trade day on after tomorrow", func(t *testing.T) {
		ex, _ := common.NewDateRange("2022-04-27", "2022-05-11")
		s.Repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex).
			Return(mockCal(), nil)
		s.Today, _ = common.NewDate("2022-04-27")
		d, err = s.GetDateByTradeDayMode("tomorrow")
		assert.NoError(t, err)
		assert.Equal(t, "2022-04-28", d.String())

		ex2, _ := common.NewDateRange("2022-04-28", "2022-05-12")
		s.Repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex2).
			Return(mockCal(), nil)
		s.Today, _ = common.NewDate("2022-04-28")
		d, err = s.GetDateByTradeDayMode("tomorrow")
		assert.NoError(t, err)
		assert.Equal(t, "2022-05-02", d.String())
	})

	t.Run("err", func(t *testing.T) {
		d, err = s.GetDateByTradeDayMode("20000101")
		assert.Error(t, err)
		assert.Nil(t, d)
		d, err = s.GetDateByTradeDayMode("aaaa")
		assert.Error(t, err)
		assert.Nil(t, d)
	})
}
