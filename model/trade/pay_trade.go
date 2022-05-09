package trade

import (
	"encoding/json"

	"github.com/aokuyama/go-stock_jp/model/order"
	"github.com/aokuyama/go-stock_jp/model/order/order_type"
	"github.com/aokuyama/go-stock_jp/model/stock"
)

type PayTrade struct {
	PayType      order_type.PayType
	SecurityCode stock.SecurityCode
	Quantity     order.Quantity
}

func NewPayTrade(trade_type string, security_code string, quantity int) (*PayTrade, error) {
	t, err := order_type.NewPayType(trade_type)
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

func (p *PayTrade) Target() *stock.SecurityCode {
	return &p.SecurityCode
}
