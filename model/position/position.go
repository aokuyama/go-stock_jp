package position

import (
	"encoding/json"
	"errors"

	"github.com/aokuyama/go-generic_subdomains/errs"
	"github.com/aokuyama/go-stock_jp/model/order/order_type"
	"github.com/aokuyama/go-stock_jp/model/stock"
	"github.com/aokuyama/go-stock_jp/model/trade"
)

type Position struct {
	positionType order_type.PositionType
	securityCode stock.SecurityCode
	quantity     Quantity
}

type positionJson struct {
	PositionType order_type.PositionType `json:"position_type"`
	SecurityCode stock.SecurityCode      `json:"security_code"`
	Quantity     Quantity                `json:"quantity"`
}

func New(position_type string, security_code string, quantity int) (*Position, error) {
	errs := errs.New()
	p, err := order_type.NewPositionType(position_type)
	errs.Append(err)
	s, err := stock.NewSecurityCode(security_code)
	errs.Append(err)
	q, err := NewQuantity(quantity)
	errs.Append(err)
	if errs.Err() != nil {
		return nil, errs.Err()
	}
	return &Position{
		positionType: *p,
		securityCode: *s,
		quantity:     *q,
	}, nil
}

func NewPositionByTrade(t *trade.PositionTrade) *Position {
	q, err := NewQuantity(t.Quantity.Int())
	if err != nil {
		panic(err)
	}
	return &Position{
		positionType: t.PositionType,
		securityCode: t.SecurityCode,
		quantity:     *q,
	}
}

func NewPositionByPositionAndPayTrade(p *Position, t *trade.PayTrade) (*Position, error) {
	if !p.IsEqualPay(t) {
		return nil, errors.New("not equal pay")
	}
	q, err := NewQuantity(p.quantity.Int() - t.Quantity.Int())
	if err != nil {
		return nil, err
	}
	return &Position{
		positionType: p.positionType,
		securityCode: t.SecurityCode,
		quantity:     *q,
	}, nil
}

func NewIntegratePosition(p1 *Position, p2 *Position) (*Position, error) {
	if !p1.IsEqualPosition(p2) {
		return nil, errors.New("not equal position")
	}
	q, err := NewQuantity(p1.quantity.Int() + p2.quantity.Int())
	if err != nil {
		return nil, err
	}
	return &Position{
		positionType: p1.positionType,
		securityCode: p2.securityCode,
		quantity:     *q,
	}, nil
}

func (p *Position) Decrement(v int) (*Position, error) {
	q, err := NewQuantity(p.Quantity() - v)
	if err != nil {
		return nil, err
	}
	return &Position{
		positionType: p.positionType,
		securityCode: p.securityCode,
		quantity:     *q,
	}, nil
}

func (p *Position) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}

func (p *Position) MarshalJSON() ([]byte, error) {
	return json.Marshal(&positionJson{
		PositionType: p.positionType,
		SecurityCode: p.securityCode,
		Quantity:     p.quantity,
	})
}

func (p *Position) UnmarshalJSON(b []byte) error {
	j := positionJson{}
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*p = Position{
		positionType: j.PositionType,
		securityCode: j.SecurityCode,
		quantity:     j.Quantity,
	}
	return nil
}

func (p *Position) IsEqualPosition(position *Position) bool {
	return p.IsEqualTarget(position) && p.positionType.IsEqual(position.Type())
}

func (p *Position) IsEqualTarget(t trade.ITrade) bool {
	return p.securityCode.IsEqual(t.Target())
}

func (p *Position) IsEqualPay(t *trade.PayTrade) bool {
	return p.IsEqualTarget(t) && p.positionType.IsCorrectPayType(&t.PayType)
}

func (p *Position) IsCompleted() bool {
	return p.quantity == 0
}

func (p *Position) Target() *stock.SecurityCode {
	return &p.securityCode
}

func (p *Position) Type() *order_type.PositionType {
	return &p.positionType
}

func (p *Position) IsEqual(b *Position) bool {
	return p.IsEqualPosition(b) && p.quantity.IsEqual(&b.quantity)
}

func (p *Position) Quantity() int {
	return p.quantity.Int()
}
