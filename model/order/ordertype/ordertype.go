package ordertype

import (
	"encoding/json"

	"github.com/aokuyama/go-generic_subdomains/errs"

	"errors"
)

type Ordertype struct {
	Trade  TradeType
	Margin *MarginType
}

func New(trade string, margin string) (*Ordertype, error) {
	errs := errs.New()
	t, err := NewTradeType(trade)
	errs.Append(err)
	var m *MarginType
	if margin == "" {
		m = nil
	} else {
		m, err = NewMarginType(margin)
		errs.Append(err)
	}
	if errs.Err() != nil {
		return nil, errs.Err()
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