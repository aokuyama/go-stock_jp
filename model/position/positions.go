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
	new := NewPositionByTrade(trade)
	for _, position := range *p {
		if position.IsEqualPosition(new) {
			return position.IntegrateIfEqualPosition(new)
		}
	}
	*p = append(*p, new)
	return nil
}

func (p *Positions) AddPayTrade(trade *trade.PayTrade) error {
	for i, position := range *p {
		if position.IsEqualPay(trade) {
			new_position, err := NewPositionByPayTrade(position, trade)
			if err != nil {
				return err
			}
			(*p)[i] = new_position
			return nil
		}
	}
	return errors.New("missing position")
}

func (p *Positions) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return (string)(j)
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

func (p *Positions) Compress() *Positions {
	return p.integrete().Uncompletes()
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
