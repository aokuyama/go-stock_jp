package service

import (
	"github.com/aokuyama/go-stock_jp/model/position"
	"github.com/aokuyama/go-stock_jp/model/trade"
)

func TradesToPositions(trades *trade.Trades) (*position.Positions, error) {
	ps := position.NewPositions()
	ps.AddPositionTrades(&trades.Positions)
	ps.AddPayTrades(&trades.Pays)
	return ps, nil
}
