package calendar

import (
	"errors"
)

type TradeDayMode string

func NewTradeDayMode(v string) (*TradeDayMode, error) {
	modes := GetTradeDayModes()
	for _, t := range modes {
		if v == t {
			c := TradeDayMode(v)
			return &c, nil
		}
	}
	return nil, errors.New("not trade day mode:" + v)
}
func GetTradeDayModes() [2]string {
	return [...]string{
		"today",
		"tomorrow",
	}
}
func (m *TradeDayMode) String() string {
	return string(*m)
}
func (m *TradeDayMode) IsTradeDayOnAfterToday() bool {
	return m.String() == "today"
}
func (m *TradeDayMode) IsTradeDayOnAfterTomorrow() bool {
	return m.String() == "tomorrow"
}
