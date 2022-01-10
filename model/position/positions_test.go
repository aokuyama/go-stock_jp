package position

import (
	"testing"

	"github.com/aokuyama/go-stock_jp/model/trade"

	"github.com/stretchr/testify/assert"
)

func TestAddPositionTrade(t *testing.T) {
	ps := NewPositions()
	assert.Equal(t, 0, len(*ps))
	t1, _ := trade.NewPositionTrade("spot_buy", "1234", 100)
	ps.AddPositionTrade(t1)
	assert.Equal(t, 1, len(*ps), "ポジションが追加される")
	f := (*ps)[0]
	t2, _ := trade.NewPositionTrade("margin_buy", "1234", 100)
	ps.AddPositionTrade(t2)
	assert.Equal(t, 2, len(*ps), "ポジションが追加される")
	ps.AddPositionTrade(t1)
	assert.Equal(t, 2, len(*ps), "ポジションがまとめられる")
	assert.Equal(t, 200, f.Quantity.Int(), "ポジションがまとめられる")
	t3, _ := trade.NewPositionTrade("spot_buy", "1234", 300)
	ps.AddPositionTrade(t3)
	ps.AddPositionTrade(t2)
	assert.Equal(t, 2, len(*ps), "ポジションがまとめられる")
	assert.Equal(t, 500, f.Quantity.Int(), "ポジションがまとめられる")
}

func TestAddPayTrade(t *testing.T) {
	var err error
	ps := NewPositions()
	po, _ := trade.NewPositionTrade("margin_buy", "1234", 200)
	err = ps.AddPositionTrade(po)
	assert.NoError(t, err)
	err = ps.AddPositionTrade(po)
	assert.NoError(t, err)
	assert.Equal(t, 400, (*ps)[0].Quantity.Int())
	py, _ := trade.NewPayTrade("pay_sell", "1234", 200)
	err = ps.AddPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 200, (*ps)[0].Quantity.Int(), "手仕舞われた分だけ減少している")
	err = ps.AddPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 0, (*ps)[0].Quantity.Int(), "手仕舞われた分だけ減少している")
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
	assert.Equal(t, 100, (*ps)[0].Quantity.Int(), "手仕舞われた分だけ減少している")
	err = ps.AddPayTrade(py)
	assert.Error(t, err, "マイナスになってしまっている")
}

func TestCompress(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		ps := NewPositions()
		p1, _ := NewPosition("spot_buy", "1234", 500)
		p2, _ := NewPosition("margin_buy", "1234", 400)
		p3, _ := NewPosition("margin_sell", "1234", 200)
		p4, _ := NewPosition("margin_sell", "1234", 100)
		p5, _ := NewPosition("margin_buy", "1235", 100)
		*ps = append(*ps, p1, p2, p3, p4, p5)
		assert.Equal(t, 5, len(*ps))
		co := ps.Compress()
		assert.Equal(t, 4, len(*co))
		assert.Equal(t, 5, len(*ps))
		assert.Equal(t, 1300, co.SumQuantity())
		assert.Equal(t, 1300, ps.SumQuantity())
		assert.Equal(t, 300, (*co)[2].Quantity.Int())
		assert.Equal(t, 200, p3.Quantity.Int())
		assert.Equal(t, 100, p4.Quantity.Int())
	})
}
func TestCopy(t *testing.T) {
	ps1 := NewPositions()
	p1, _ := NewPosition("spot_buy", "1234", 500)
	p2, _ := NewPosition("margin_buy", "1234", 400)
	*ps1 = append(*ps1, p1, p2)
	ps2 := ps1.Copy()
	assert.Equal(t, 900, ps1.SumQuantity())
	assert.Equal(t, 900, ps2.SumQuantity())
	(*ps2)[0].Quantity -= 400
	assert.Equal(t, 500, ps2.SumQuantity())
	assert.Equal(t, 900, ps1.SumQuantity(), "コピー元に影響がない")
}
