package trade

import (
	"encoding/json"

	"github.com/aokuyama/go-stock_jp/model/order"
	"github.com/aokuyama/go-stock_jp/model/order/order_type"
	"github.com/aokuyama/go-stock_jp/model/stock"
)

type PositionTrade struct {
	PositionType order_type.PositionType
	SecurityCode stock.SecurityCode
	Quantity     order.Quantity
}

func NewPositionTrade(trade_type string, security_code string, quantity int) (*PositionTrade, error) {
	t, err := order_type.NewPositionType(trade_type)
	if err != nil {
		return nil, err
	}
	s, err2 := stock.NewSecurityCode(security_code)
	if err2 != nil {
		return nil, err2
	}
	q, err3 := order.NewQuantity(quantity)
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

func (p *PositionTrade) Target() *stock.SecurityCode {
	return &p.SecurityCode
}

func (p *PositionTrade) Type() *order_type.PositionType {
	return &p.PositionType
}
