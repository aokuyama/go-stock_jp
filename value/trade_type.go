package value

import (
	"errors"
)

type TradeType interface {
	String() string
	IsPosition() bool
	IsPay() bool
}

func NewTradeType(v string) (TradeType, error) {
	var tt TradeType
	var err error
	tt, err = NewPositionType(v)
	if err == nil {
		return tt, nil
	}
	tt, err = NewPayType(v)
	if err == nil {
		return tt, nil
	}
	return nil, errors.New("Invalid trade type:" + v)
}
