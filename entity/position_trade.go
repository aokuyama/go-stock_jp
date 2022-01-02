package entity

import (
	"encoding/json"
	"github.com/aokuyama/go-stock_jp/value"
)

type PositionTrade struct {
	PositionType value.PositionType
	SecurityCode value.SecurityCode
	Quantity     value.Quantity
}

func NewPositionTrade(trade_type string, security_code string, quantity int) (*PositionTrade, error) {
	t, err := value.NewPositionType(trade_type)
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
	return &PositionTrade{
		PositionType: *t,
		SecurityCode: *s,
		Quantity:     *q,
	}, nil
}

func (p *PositionTrade) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}

func (p *PositionTrade) Target() *value.SecurityCode {
	return &p.SecurityCode
}

func (p *PositionTrade) Type() *value.PositionType {
	return &p.PositionType
}
