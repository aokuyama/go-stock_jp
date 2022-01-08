package position

import (
	"encoding/json"
	"errors"

	"github.com/aokuyama/go-stock_jp/model/order"
	"github.com/aokuyama/go-stock_jp/model/stock"
	"github.com/aokuyama/go-stock_jp/model/trade"
)

type Position struct {
	PositionType order.PositionType `json:"position_type"`
	SecurityCode stock.SecurityCode `json:"security_code"`
	Quantity     order.Quantity     `json:"quantity"`
}

func NewPosition(position_type string, security_code string, quantity int) (*Position, error) {
	p, err := order.NewPositionType(position_type)
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
	return &Position{
		PositionType: *p,
		SecurityCode: *s,
		Quantity:     *q,
	}, nil
}

func NewErrorPosition(trade *trade.PayTrade) (*Position, error) {
	p, err := trade.PayType.PairPositionType()
	if err != nil {
		return nil, err
	}
	q, err := order.NewErrorQuantity(trade.Quantity.Int())
	if err != nil {
		return nil, err
	}
	return &Position{
		PositionType: *p,
		SecurityCode: trade.SecurityCode,
		Quantity:     *q,
	}, nil
}

func (p *Position) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}

func (p *Position) IsEqualPosition(t trade.IPosition) bool {
	return p.IsEqualTarget(t) && p.PositionType.IsEqual(t.Type())
}

func (p *Position) IsEqualTarget(t trade.ITrade) bool {
	return p.SecurityCode.IsEqual(t.Target())
}

func (p *Position) IncludePosition(t *trade.PositionTrade) error {
	if !p.IsEqualPosition(t) {
		return errors.New("not equal position")
	}
	p.Quantity += t.Quantity
	return nil
}

func (p *Position) IsEqualPay(t *trade.PayTrade) bool {
	return p.IsEqualTarget(t) && p.PositionType.IsCorrectPayType(&t.PayType)
}

func (p *Position) IncludePay(t *trade.PayTrade) error {
	if !p.IsEqualPay(t) {
		return errors.New("not equal pay")
	}
	p.Quantity -= t.Quantity
	return nil
}

func (p *Position) IsError() bool {
	return p.Quantity.IsError()
}

func (p *Position) IsCompleted() bool {
	return p.Quantity == 0
}

func (p *Position) Target() *stock.SecurityCode {
	return &p.SecurityCode
}

func (p *Position) Type() *order.PositionType {
	return &p.PositionType
}

func (p *Position) integrate(position *Position) error {
	if p == position {
		return errors.New("same instance")
	}
	if !p.IsEqualPosition(position) {
		return errors.New("not equal position")
	}
	p.Quantity += position.Quantity
	position.Quantity = 0
	return nil
}

func (p *Position) IntegrateIfEqualPosition(position *Position) error {
	if !p.IsEqualPosition(position) {
		return nil
	}
	return p.integrate(position)
}

func (p *Position) offset(position *Position) error {
	if p == position {
		return errors.New("same instance")
	}
	if !p.IsEqualPosition(position) {
		return errors.New("not equal position")
	}
	p.Quantity -= position.Quantity
	position.Quantity = 0
	return nil
}
func (p *Position) OffsetIfEqualPosition(position *Position) error {
	if !p.IsEqualPosition(position) {
		return nil
	}
	return p.offset(position)
}
