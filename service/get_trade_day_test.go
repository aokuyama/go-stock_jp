package service

import (
	"testing"

	"github.com/aokuyama/go-stock_jp/common"
	"github.com/aokuyama/go-stock_jp/model/calendar"
	"github.com/aokuyama/go-stock_jp/model/calendar/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func getNewTestGetTradeDayService(t *testing.T) *GetTradeDayService {
	gomock.NewController(t)
	mc := gomock.NewController(t)
	defer mc.Finish()
	repo := mock.NewMockCalendarRepository(mc)
	return NewGetTradeDayService(repo)
}

func TestGetTradeDay(t *testing.T) {
	s := getNewTestGetTradeDayService(t)
	ex, _ := common.NewDateRange("2022-01-16", "2022-01-30")
	s.repository.(*mock.MockCalendarRepository).EXPECT().
		LoadByDateRange(ex).
		Return(newCalendar(), nil)
	t.Run("after tomorrow", func(t *testing.T) {
		r, _ := s.GetTradeDayOnAfterTomorrow("2022-01-16")
		assert.Equal(t, "2022-01-18", r.String())
	})
	s.repository.(*mock.MockCalendarRepository).EXPECT().
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
	c, _ := calendar.NewCalendar(&dates)
	return c
}

func TestErrorGetTradeDayOnAfterTomorrow(t *testing.T) {
	s := getNewTestGetTradeDayService(t)
	r, err := s.GetTradeDayOnAfterTomorrow("あ")
	assert.Error(t, err)
	assert.Nil(t, r)
	r2, err2 := s.GetTradeDayOnAfterToday("")
	assert.Error(t, err2)
	assert.Nil(t, r2)
	t.Run("not found", func(t *testing.T) {
		ex, _ := common.NewDateRange("2022-01-19", "2022-02-02")
		s.repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex).
			Return(newCalendar(), nil)
		r, err := s.GetTradeDayOnAfterTomorrow("2022-01-19")
		assert.Error(t, err)
		assert.Nil(t, r)
	})
	t.Run("not found", func(t *testing.T) {
		ex, _ := common.NewDateRange("2022-01-19", "2022-02-02")
		s.repository.(*mock.MockCalendarRepository).EXPECT().
			LoadByDateRange(ex).
			Return(newCalendar(), nil)
		r, err := s.GetTradeDayOnAfterToday("2022-01-19")
		assert.Error(t, err)
		assert.Nil(t, r)
	})
}
