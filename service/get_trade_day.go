package service

import (
	"errors"
	"time"

	"github.com/aokuyama/go-stock_jp/model/calendar"
	"github.com/aokuyama/go-stock_jp/model/common"
)

type GetTradeDay struct {
	Repository calendar.CalendarRepository
	Today      *common.Date
}

func NewGetTradeDay(repository calendar.CalendarRepository) (*GetTradeDay, error) {
	now := time.Now()
	today, err := common.NewDateByTime(&now)
	if err != nil {
		return nil, err
	}
	return &GetTradeDay{
		Repository: repository,
		Today:      today,
	}, nil
}

func (s *GetTradeDay) GetDateByTradeDayMode(mode string) (*common.Date, error) {
	m, _ := calendar.NewTradeDayMode(mode)
	if m == nil {
		return common.NewDate(mode)
	}
	if m.IsTradeDayOnAfterToday() {
		return s.GetTradeDayOnAfterToday(s.Today.String())
	}
	if m.IsTradeDayOnAfterTomorrow() {
		return s.GetTradeDayOnAfterTomorrow(s.Today.String())
	}
	return nil, errors.New("fatal error")
}

func (s *GetTradeDay) getCalendar(begin *common.Date) (*calendar.Calendar, error) {
	// 2週間分くらいあれば良い
	end := common.NewDateAdded(begin, time.Hour*24*14)
	r, err := common.NewDateRangeByDates(begin, end)
	if err != nil {
		return nil, err
	}
	return s.Repository.LoadByDateRange(r)
}

func (s *GetTradeDay) GetTradeDayOnAfterTomorrow(today string) (*common.Date, error) {
	d, err := common.NewDate(today)
	if err != nil {
		return nil, err
	}
	calendar, err := s.getCalendar(d)
	if err != nil {
		return nil, err
	}
	r := calendar.TradeDayOnAfterTomorrow(d)
	if r == nil {
		return nil, errors.New("not found trade day")
	}
	return r, nil
}

func (s *GetTradeDay) GetTradeDayOnAfterToday(today string) (*common.Date, error) {
	d, err := common.NewDate(today)
	if err != nil {
		return nil, err
	}
	calendar, err := s.getCalendar(d)
	if err != nil {
		return nil, err
	}
	r := calendar.TradeDayOnAfterToday(d)
	if r == nil {
		return nil, errors.New("not found trade day")
	}
	return r, nil
}
