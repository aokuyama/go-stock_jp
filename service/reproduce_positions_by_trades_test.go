package service

import (
	"testing"

	"github.com/aokuyama/go-stock_jp/model/position"
	"github.com/aokuyama/go-stock_jp/model/trade"

	"github.com/stretchr/testify/assert"
)

func TestAddPositionTrades(t *testing.T) {
	ts := trade.NewTrades()
	ps := position.NewPositions()
	addPositionTrades(ps, &ts.Positions)
	assert.Equal(t, 0, len(*ps))
	ts.AddNewTrade("spot_buy", "1234", 100)
	ts.AddNewTrade("spot_buy", "1235", 100)
	ts.AddNewTrade("margin_buy", "1235", 100)
	addPositionTrades(ps, &ts.Positions)
	assert.Equal(t, 3, len(*ps), "ポジションが生成される")
}
func TestCollectPositions(t *testing.T) {
	ts := trade.NewTrades()
	ps := position.NewPositions()
	ts.AddNewTrade("spot_buy", "1234", 100)
	ts.AddNewTrade("spot_buy", "1234", 100)
	addPositionTrades(ps, &ts.Positions)
	assert.Equal(t, 1, len(*ps), "同じポジションはまとめられる")
	assert.Equal(t, 200, (*ps)[0].Quantity.Int(), "同じポジションはまとめられる")
}
func TestImportAllTrade(t *testing.T) {
	ts := trade.NewTrades()
	ps := position.NewPositions()
	assert.Equal(t, 0, len(*ps))
	ts.AddNewTrade("spot_buy", "1234", 1000)
	ts.AddNewTrade("margin_buy", "1234", 200)
	ts.AddNewTrade("margin_sell", "1234", 500)
	ts.AddNewTrade("margin_sell", "1234", 300)
	ts.AddNewTrade("margin_buy", "2534", 200)

	ts.AddNewTrade("pay_sell", "1234", 100)
	ts.AddNewTrade("pay_buy", "1234", 800)
	ts.AddNewTrade("pay_sell", "1234", 100)
	ts.AddNewTrade("spot_sell", "1234", 1000)
	ts.AddNewTrade("pay_sell", "2534", 200)
	addPositionTrades(ps, &ts.Positions)
	assert.Equal(t, 4, len(*ps.Uncompletes()))
	addPayTrades(ps, &ts.Pays)
	assert.Equal(t, 4, len(*ps), "ポジションが生成される")
	assert.Equal(t, 0, len(*ps.Uncompletes()), "未完成なし")
}
