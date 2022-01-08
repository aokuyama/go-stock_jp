package trade

import (
	"errors"
)

type PositionTrades []*PositionTrade
type PayTrades []*PayTrade

type Trades struct {
	Positions PositionTrades
	Pays      PayTrades
}

func NewTrades() *Trades {
	return &Trades{}
}

func (t *Trades) AddNewTrade(trade_type string, security_code string, quantity int) (ITrade, error) {
	pos, err := NewPositionTrade(trade_type, security_code, quantity)
	if err == nil {
		t.Positions = append(t.Positions, pos)
		return pos, nil
	}
	pay, err := NewPayTrade(trade_type, security_code, quantity)
	if err == nil {
		t.Pays = append(t.Pays, pay)
		return pay, nil
	}
	return nil, errors.New("failed parse trade")
}
