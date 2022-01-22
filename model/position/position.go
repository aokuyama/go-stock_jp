package position

import (
	"encoding/json"
	"errors"

	"github.com/aokuyama/go-stock_jp/model/order/ordertype"
	"github.com/aokuyama/go-stock_jp/model/stock"
	"github.com/aokuyama/go-stock_jp/model/trade"
)

type Position struct {
	PositionType ordertype.PositionType `json:"position_type"`
	SecurityCode stock.SecurityCode     `json:"security_code"`
	Quantity     Quantity               `json:"quantity"`
}

func NewPosition(position_type string, security_code string, quantity int) (*Position, error) {
	var err error
	p, err := ordertype.NewPositionType(position_type)
	if err != nil {
		return nil, err
	}
	s, err := stock.NewSecurityCode(security_code)
	if err != nil {
		return nil, err
	}
	q, err := NewQuantity(quantity)
	if err != nil {
		return nil, err
	}
	return &Position{
		PositionType: *p,
		SecurityCode: *s,
		Quantity:     *q,
	}, nil
}

func NewPositionByTrade(t *trade.PositionTrade) *Position {
	q, err := NewQuantity(t.Quantity.Int())
	if err != nil {
		panic(err)
	}
	return &Position{
		PositionType: t.PositionType,
		SecurityCode: t.SecurityCode,
		Quantity:     *q,
	}
}

func NewPositionByPositionAndPayTrade(p *Position, t *trade.PayTrade) (*Position, error) {
	if !p.IsEqualPay(t) {
		return nil, errors.New("not equal pay")
	}
	q, err := NewQuantity(p.Quantity.Int() - t.Quantity.Int())
	if err != nil {
		return nil, err
	}
	return &Position{
		PositionType: p.PositionType,
		SecurityCode: t.SecurityCode,
		Quantity:     *q,
	}, nil
}

func NewIntegratePosition(p1 *Position, p2 *Position) (*Position, error) {
	if !p1.IsEqualPosition(p2) {
		return nil, errors.New("not equal position")
	}
	q, err := NewQuantity(p1.Quantity.Int() + p2.Quantity.Int())
	if err != nil {
		return nil, err
	}
	return &Position{
		PositionType: p1.PositionType,
		SecurityCode: p2.SecurityCode,
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

func (p *Position) IsEqualPosition(position *Position) bool {
	return p.IsEqualTarget(position) && p.PositionType.IsEqual(position.Type())
}

func (p *Position) IsEqualTarget(t trade.ITrade) bool {
	return p.SecurityCode.IsEqual(t.Target())
}

func (p *Position) IsEqualPay(t *trade.PayTrade) bool {
	return p.IsEqualTarget(t) && p.PositionType.IsCorrectPayType(&t.PayType)
}

func (p *Position) IsCompleted() bool {
	return p.Quantity == 0
}

func (p *Position) Target() *stock.SecurityCode {
	return &p.SecurityCode
}

func (p *Position) Type() *ordertype.PositionType {
	return &p.PositionType
}

func (p *Position) IsEqual(b *Position) bool {
	return p.IsEqualPosition(b) && p.Quantity.IsEqual(&b.Quantity)
}
