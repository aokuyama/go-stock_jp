package position_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/position"
	"github.com/aokuyama/go-stock_jp/model/trade"

	"github.com/stretchr/testify/assert"
)

func TestAddPositionTrade(t *testing.T) {
	ps := NewPositions()
	assert.Equal(t, 0, len(*ps))
	t1, _ := trade.NewPositionTrade("spot_buy", "1234", 100)
	ps.AddPositionTrade(t1)
	assert.Equal(t, 1, len(*ps), "ポジションが追加される")
	t2, _ := trade.NewPositionTrade("margin_buy", "1234", 100)
	ps.AddPositionTrade(t2)
	assert.Equal(t, 2, len(*ps), "ポジションが追加される")
	ps.AddPositionTrade(t1)
	assert.Equal(t, 2, len(*ps), "ポジションがまとめられる")
	assert.Equal(t, 200, (*ps)[0].Quantity(), "ポジションがまとめられる")
	t3, _ := trade.NewPositionTrade("spot_buy", "1234", 300)
	ps.AddPositionTrade(t3)
	ps.AddPositionTrade(t2)
	assert.Equal(t, 2, len(*ps), "ポジションがまとめられる")
	assert.Equal(t, 500, (*ps)[0].Quantity(), "ポジションがまとめられる")
}

func TestAddPayTrade(t *testing.T) {
	var err error
	ps := NewPositions()
	po, _ := trade.NewPositionTrade("margin_buy", "1234", 200)
	err = ps.AddPositionTrade(po)
	assert.NoError(t, err)
	err = ps.AddPositionTrade(po)
	assert.NoError(t, err)
	assert.Equal(t, 400, (*ps)[0].Quantity())
	py, _ := trade.NewPayTrade("pay_sell", "1234", 200)
	err = ps.AddPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 200, (*ps)[0].Quantity(), "手仕舞われた分だけ減少している")
	err = ps.AddPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 0, (*ps)[0].Quantity(), "手仕舞われた分だけ減少している")
}

func TestErrorPayTrade(t *testing.T) {
	var err error
	ps := NewPositions()
	py, _ := trade.NewPayTrade("spot_sell", "1234", 100)
	err = ps.AddPayTrade(py)
	assert.Error(t, err, "対象がないためエラー")
	po, _ := trade.NewPositionTrade("margin_buy", "1234", 100)
	ps.AddPositionTrade(po)
	py2, _ := trade.NewPayTrade("pay_sell", "2234", 100)
	err = ps.AddPayTrade(py2)
	assert.Error(t, err, "対象がないためエラー")
}
func TestOverPayTrade(t *testing.T) {
	var err error
	ps := NewPositions()
	po, _ := trade.NewPositionTrade("margin_buy", "1234", 300)
	ps.AddPositionTrade(po)
	py, _ := trade.NewPayTrade("pay_sell", "1234", 200)
	err = ps.AddPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 100, (*ps)[0].Quantity(), "手仕舞われた分だけ減少している")
	err = ps.AddPayTrade(py)
	assert.Error(t, err, "マイナスになってしまっている")
}
