package service

import (
	"errors"
	"time"

	"github.com/aokuyama/go-stock_jp/common"
	"github.com/aokuyama/go-stock_jp/model/calendar"
)

type GetTradeDayService struct {
	Repository calendar.CalendarRepository
}

func NewGetTradeDayService(repository calendar.CalendarRepository) *GetTradeDayService {
	return &GetTradeDayService{
		Repository: repository,
	}
}

func (s *GetTradeDayService) getCalendar(begin *common.Date) (*calendar.Calendar, error) {

	// 2週間分くらいあれば良い
	end := common.NewDateAdded(begin, time.Hour*24*14)
	r, err := common.NewDateRangeByDates(begin, end)
	if err != nil {
		return nil, err
	}
	return s.Repository.LoadByDateRange(r)
}

func (s *GetTradeDayService) GetTradeDayOnAfterTomorrow(today string) (*common.Date, error) {
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

func (s *GetTradeDayService) GetTradeDayOnAfterToday(today string) (*common.Date, error) {
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
