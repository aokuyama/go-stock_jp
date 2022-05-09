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

func (p *Positions) AddPosition(position *Position) error {
	for i, existing := range *p {
		if existing.IsEqualPosition(position) {
			ip, err := NewIntegratePosition(existing, position)
			if err != nil {
				return err
			}
			(*p)[i] = ip
			return nil
		}
	}
	*p = append(*p, position)
	return nil
}

func (p *Positions) AddPositionTrade(trade *trade.PositionTrade) error {
	return p.AddPosition(NewPositionByTrade(trade))
}

func (p *Positions) AddPayTrade(trade *trade.PayTrade) error {
	for i, position := range *p {
		if position.IsEqualPay(trade) {
			new_position, err := NewPositionByPositionAndPayTrade(position, trade)
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

func (p *Positions) SumQuantity() int {
	sum := 0
	for _, pos := range *p {
		sum += pos.quantity.Int()
	}
	return sum
}
