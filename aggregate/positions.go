package aggregate

import (
	"encoding/json"
	"errors"
	"github.com/aokuyama/go-stock_jp/entity"
)

type Positions []*entity.Position

func NewPositions() *Positions {
	return &Positions{}
}

func (p *Positions) addPositionTrade(trade *entity.PositionTrade) error {
	for _, position := range *p {
		if position.IsEqualPosition(trade) {
			return position.IncludePosition(trade)
		}
	}
	new_pos := entity.Position{
		PositionType: trade.PositionType,
		SecurityCode: trade.SecurityCode,
		Quantity:     trade.Quantity,
	}
	*p = append(*p, &new_pos)
	return nil
}

func (p *Positions) addPayTrade(trade *entity.PayTrade) error {
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
	new_pos, err := entity.NewErrorPosition(trade)
	if err != nil {
		panic(err)
	}
	*p = append(*p, new_pos)
	return errors.New("missing position")
}

func (p *Positions) errors() *Positions {
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

func (p *Positions) addPositionTrades(trades *PositionTrades) error {
	for _, trade := range *trades {
		err := p.addPositionTrade(trade)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Positions) addPayTrades(trades *PayTrades) error {
	for _, trade := range *trades {
		err := p.addPayTrade(trade)
		if err != nil {
			return err
		}
	}
	return nil
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
	copy := p.copy()
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

func (p *Positions) Diff(dests *Positions) *Positions {
	copy_src := p.copy()
	copy_dests := dests.copy()
	for _, dest := range *copy_dests {
		for _, src := range *copy_src {
			src.OffsetIfEqualPosition(dest)
		}
	}
	for _, dest := range *copy_dests {
		if !dest.IsCompleted() {
			dest.Quantity *= -1
			*copy_src = append(*copy_src, dest)
		}
	}
	return copy_src.Compress()
}

func (p *Positions) sumQuantity() int {
	sum := 0
	for _, pos := range *p {
		sum += pos.Quantity.Int()
	}
	return sum
}
func (p *Positions) copy() *Positions {
	copy := NewPositions()
	for _, pos := range *p {
		cp := entity.Position{}
		cp = *pos
		*copy = append(*copy, &cp)
	}
	return copy
}
