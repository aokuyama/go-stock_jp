package ordertype

import (
	"encoding/json"
	"errors"
)

type Ordertype struct {
	Trade  TradeType
	Margin *MarginType
}

func New(trade string, margin string) (*Ordertype, error) {
	t, err := NewTradeType(trade)
	if err != nil {
		return nil, err
	}
	var m *MarginType
	if margin == "" {
		m = nil
	} else {
		m, err = NewMarginType(margin)
	}
	if err != nil {
		return nil, err
	}
	ot := Ordertype{
		t,
		m,
	}
	if ot.Error() != "" {
		return nil, errors.New(ot.Error())
	}
	return &ot, nil
}

func (t *Ordertype) Error() string {
	if t.Trade.IsSpot() {
		if t.Margin != nil {
			return "invalid order type:" + t.String()
		}
	} else {
		if t.Margin == nil {
			return "invalid order type:" + t.String()
		}
	}
	return ""
}

func (t *Ordertype) String() string {
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}
