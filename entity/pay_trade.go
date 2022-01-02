package entity

import (
	"encoding/json"
	"github.com/aokuyama/go-stock_jp/value"
)

type PayTrade struct {
	PayType      value.PayType
	SecurityCode value.SecurityCode
	Quantity     value.Quantity
}

func NewPayTrade(trade_type string, security_code string, quantity int) (*PayTrade, error) {
	t, err := value.NewPayType(trade_type)
	if err != nil {
		return nil, err
	}
	s, err2 := value.NewSecurityCode(security_code)
	if err2 != nil {
		return nil, err2
	}
	q, err3 := value.NewQuantity(quantity)
	if err3 != nil {
		return nil, err3
	}
	return &PayTrade{
		PayType:      *t,
		SecurityCode: *s,
		Quantity:     *q,
	}, nil
}

func (p *PayTrade) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}

func (p *PayTrade) Target() *value.SecurityCode {
	return &p.SecurityCode
}
