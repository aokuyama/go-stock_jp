package service

import (
	"github.com/aokuyama/go-stock_jp/model/position"
	"github.com/aokuyama/go-stock_jp/model/trade"
)

func ReproducePositionsByTrades(trades *trade.Trades) (*position.Positions, error) {
	var err error
	ps := position.NewPositions()
	err = AddPositionTrades(ps, &trades.Positions)
	if err != nil {
		return nil, err
	}
	err = AddPayTrades(ps, &trades.Pays)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func AddPositionTrades(positions *position.Positions, trades *trade.PositionTrades) error {
	for _, trade := range *trades {
		err := positions.AddPositionTrade(trade)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddPayTrades(positions *position.Positions, trades *trade.PayTrades) error {
	for _, trade := range *trades {
		err := positions.AddPayTrade(trade)
		if err != nil {
			return err
		}
	}
	return nil
}
