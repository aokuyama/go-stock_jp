package aggregate

import (
	"errors"
	"github.com/aokuyama/go-stock_jp/entity"
)

type PositionTrades []*entity.PositionTrade
type PayTrades []*entity.PayTrade

type Trades struct {
	Positions PositionTrades
	Pays      PayTrades
}

func NewTrades() *Trades {
	return &Trades{}
}

func (t *Trades) AddNewTrade(trade_type string, security_code string, quantity int) (entity.ITrade, error) {
	pos, err := entity.NewPositionTrade(trade_type, security_code, quantity)
	if err == nil {
		t.Positions = append(t.Positions, pos)
		return pos, nil
	}
	pay, err := entity.NewPayTrade(trade_type, security_code, quantity)
	if err == nil {
		t.Pays = append(t.Pays, pay)
		return pay, nil
	}
	return nil, errors.New("failed parse trade")
}

func (t *Trades) ToPositions() (*Positions, error) {
	ps := NewPositions()
	ps.addPositionTrades(&t.Positions)
	ps.addPayTrades(&t.Pays)
	return ps, nil
}
