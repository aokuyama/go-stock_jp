package position

import (
	"encoding/json"
	"errors"

	"github.com/aokuyama/go-stock_jp/model/trade"
)

type Positions []*Position

func NewPositions() *Positions {
	return &Positions{}
}

func (p *Positions) AddPositionTrade(trade *trade.PositionTrade) error {
	for _, position := range *p {
		if position.IsEqualPosition(trade) {
			return position.IncludePosition(trade)
		}
	}
	new_pos := Position{
		PositionType: trade.PositionType,
		SecurityCode: trade.SecurityCode,
		Quantity:     trade.Quantity,
	}
	*p = append(*p, &new_pos)
	return nil
}

func (p *Positions) AddPayTrade(trade *trade.PayTrade) error {
	for _, position := range *p {
		if position.IsEqualPay(trade) {
			err := position.IncludePay(trade)
			if err != nil {
				return err
			}
			if position.IsError() {
				return errors.New("over payment")
			}
			return nil
		}
	}
	new_pos, err := NewErrorPosition(trade)
	if err != nil {
		panic(err)
	}
	*p = append(*p, new_pos)
	return errors.New("missing position")
}

func (p *Positions) Errors() *Positions {
	ps := NewPositions()
	for _, position := range *p {
		if position.IsError() {
			*ps = append(*ps, position)
		}
	}
	return ps
}

func (p *Positions) Uncompletes() *Positions {
	ps := NewPositions()
	for _, position := range *p {
		if !position.IsCompleted() {
			*ps = append(*ps, position)
		}
	}
	return ps
}

func (p *Positions) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}

func (p *Positions) Compress() *Positions {
	return p.integrete().filter()
}

func (p *Positions) integrete() *Positions {
	copy := p.Copy()
	for _, a := range *copy {
		for _, b := range *copy {
			if a == b {
				continue
			}
			err := a.IntegrateIfEqualPosition(b)
			if err != nil {
				panic(err)
			}
		}
	}
	return copy
}

func (p *Positions) filter() *Positions {
	ps := NewPositions()
	for _, pos := range *p {
		if !pos.IsCompleted() {
			*ps = append(*ps, pos)
		}
	}
	return ps
}

func (p *Positions) SumQuantity() int {
	sum := 0
	for _, pos := range *p {
		sum += pos.Quantity.Int()
	}
	return sum
}
func (p *Positions) Copy() *Positions {
	copy := NewPositions()
	for _, pos := range *p {
		cp := Position{}
		cp = *pos
		*copy = append(*copy, &cp)
	}
	return copy
}
