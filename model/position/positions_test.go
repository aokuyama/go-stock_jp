package position

import (
	"testing"

	"github.com/aokuyama/go-stock_jp/model/trade"

	"github.com/stretchr/testify/assert"
)

func TestPickNewPosition(t *testing.T) {
	ps := NewPositions()
	assert.Equal(t, 0, len(*ps))
	t1, _ := trade.NewPositionTrade("spot_buy", "1234", 100)
	ps.addPositionTrade(t1)
	assert.Equal(t, 1, len(*ps), "ポジションが追加される")
	f := (*ps)[0]
	t2, _ := trade.NewPositionTrade("margin_buy", "1234", 100)
	ps.addPositionTrade(t2)
	assert.Equal(t, 2, len(*ps), "ポジションが追加される")
	ps.addPositionTrade(t1)
	assert.Equal(t, 2, len(*ps), "ポジションがまとめられる")
	assert.Equal(t, 200, f.Quantity.Int(), "ポジションがまとめられる")
	t3, _ := trade.NewPositionTrade("spot_buy", "1234", 300)
	ps.addPositionTrade(t3)
	ps.addPositionTrade(t2)
	assert.Equal(t, 2, len(*ps), "ポジションがまとめられる")
	assert.Equal(t, 500, f.Quantity.Int(), "ポジションがまとめられる")
}
func TestAddPositionTrades(t *testing.T) {
	ts := trade.NewTrades()
	ps := NewPositions()
	ps.AddPositionTrades(&ts.Positions)
	assert.Equal(t, 0, len(*ps))
	ts.AddNewTrade("spot_buy", "1234", 100)
	ts.AddNewTrade("spot_buy", "1235", 100)
	ts.AddNewTrade("margin_buy", "1235", 100)
	ps.AddPositionTrades(&ts.Positions)
	assert.Equal(t, 3, len(*ps), "ポジションが生成される")
}
func TestCollectPositions(t *testing.T) {
	ts := trade.NewTrades()
	ps := NewPositions()
	ts.AddNewTrade("spot_buy", "1234", 100)
	ts.AddNewTrade("spot_buy", "1234", 100)
	ps.AddPositionTrades(&ts.Positions)
	assert.Equal(t, 1, len(*ps), "同じポジションはまとめられる")
	assert.Equal(t, 200, (*ps)[0].Quantity.Int(), "同じポジションはまとめられる")
}
func TestAddPayTrade(t *testing.T) {
	var err error
	ps := NewPositions()
	po, _ := trade.NewPositionTrade("margin_buy", "1234", 200)
	err = ps.addPositionTrade(po)
	assert.NoError(t, err)
	err = ps.addPositionTrade(po)
	assert.NoError(t, err)
	assert.Equal(t, 400, (*ps)[0].Quantity.Int())
	py, _ := trade.NewPayTrade("pay_sell", "1234", 200)
	err = ps.addPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 200, (*ps)[0].Quantity.Int(), "手仕舞われた分だけ減少している")
	err = ps.addPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 0, (*ps)[0].Quantity.Int(), "手仕舞われた分だけ減少している")
}

func TestErrorPayTrade(t *testing.T) {
	var err error
	ps := NewPositions()
	assert.Equal(t, 0, len(*ps.errors()))
	py, _ := trade.NewPayTrade("spot_sell", "1234", 100)
	err = ps.addPayTrade(py)
	assert.Error(t, err)
	assert.Equal(t, 1, len(*ps.errors()), "対象がないためエラーが追加される")
	ep := (*ps.errors())[0]
	assert.Equal(t, -100, ep.Quantity.Int())
	po, _ := trade.NewPositionTrade("margin_buy", "1234", 100)
	ps.addPositionTrade(po)
	py2, _ := trade.NewPayTrade("pay_sell", "2234", 100)
	err = ps.addPayTrade(py2)
	assert.Error(t, err)
	assert.Equal(t, 2, len(*ps.errors()), "対象がないためエラーが追加される")
	err = ps.addPayTrade(py)
	assert.Error(t, err, "エラーポジションへ更に追加")
	assert.Equal(t, -200, ep.Quantity.Int())
	assert.Equal(t, 2, len(*ps.errors()), "エラーとはいえ対象があったため、エラーが追加されない")
}

func TestOverPayTrade(t *testing.T) {
	var err error
	ps := NewPositions()
	po, _ := trade.NewPositionTrade("margin_buy", "1234", 300)
	ps.addPositionTrade(po)
	py, _ := trade.NewPayTrade("pay_sell", "1234", 200)
	err = ps.addPayTrade(py)
	assert.NoError(t, err)
	assert.Equal(t, 100, (*ps)[0].Quantity.Int(), "手仕舞われた分だけ減少している")
	err = ps.addPayTrade(py)
	assert.Error(t, err)
	assert.Equal(t, -100, (*ps)[0].Quantity.Int(), "マイナスになってしまっている")
}

func TestImportAllTrade(t *testing.T) {
	ts := trade.NewTrades()
	ps := NewPositions()
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
	ps.AddPositionTrades(&ts.Positions)
	assert.Equal(t, 4, len(*ps.Uncompletes()))
	ps.AddPayTrades(&ts.Pays)
	assert.Equal(t, 4, len(*ps), "ポジションが生成される")
	assert.Equal(t, 0, len(*ps.errors()), "エラーなし")
	assert.Equal(t, 0, len(*ps.Uncompletes()), "未完成なし")
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
		assert.Equal(t, 1300, co.sumQuantity())
		assert.Equal(t, 1300, ps.sumQuantity())
		assert.Equal(t, 300, (*co)[2].Quantity.Int())
		assert.Equal(t, 200, p3.Quantity.Int())
		assert.Equal(t, 100, p4.Quantity.Int())
	})
}
func TestDiff(t *testing.T) {
	t.Run("no diff", func(t *testing.T) {
		ps1 := NewPositions()
		p1, _ := NewPosition("spot_buy", "1234", 500)
		p2, _ := NewPosition("margin_buy", "1234", 400)
		p3, _ := NewPosition("margin_sell", "1234", 200)
		p4, _ := NewPosition("margin_sell", "1234", 100)
		p5, _ := NewPosition("margin_buy", "1235", 100)
		*ps1 = append(*ps1, p1, p2, p3, p4, p5)
		ps2 := NewPositions()
		*ps2 = append(*ps2, p5, p4, p3, p2, p1)
		diff1 := ps1.Diff(ps2)
		assert.Equal(t, 0, len(*diff1), diff1.String())
		diff2 := ps2.Diff(ps1)
		assert.Equal(t, 0, len(*diff2), diff2.String())
	})
	t.Run("case 1", func(t *testing.T) {
		ps1 := NewPositions()
		p1, _ := NewPosition("spot_buy", "1234", 500)
		p2, _ := NewPosition("margin_buy", "1234", 400)
		p3, _ := NewPosition("margin_sell", "1234", 200)
		p4, _ := NewPosition("margin_sell", "1234", 100)
		p5, _ := NewPosition("margin_buy", "1235", 100)
		*ps1 = append(*ps1, p1, p2, p3, p4, p5)
		ps2 := NewPositions()
		*ps2 = append(*ps2, p1, p2, p3, p5)
		diff1 := ps1.Diff(ps2)
		assert.Equal(t, 1, len(*diff1))
		assert.Equal(t, 100, (*diff1)[0].Quantity.Int())
		diff2 := ps2.Diff(ps1)
		assert.Equal(t, 1, len(*diff2))
		assert.Equal(t, -100, (*diff2)[0].Quantity.Int())
		assert.Equal(t, 5, len(*ps1))
		assert.Equal(t, 1300, ps1.sumQuantity())
		assert.Equal(t, 1200, ps2.sumQuantity())
	})
	t.Run("case 2", func(t *testing.T) {
		ps1 := NewPositions()
		p1, _ := NewPosition("spot_buy", "1234", 500)
		p2, _ := NewPosition("margin_buy", "1234", 400)
		p3, _ := NewPosition("margin_sell", "1234", 300)
		p4, _ := NewPosition("margin_buy", "1235", 200)
		*ps1 = append(*ps1, p1, p2, p3)
		ps2 := NewPositions()
		*ps2 = append(*ps2, p2, p3, p4)
		diff1 := ps1.Diff(ps2)
		assert.Equal(t, 2, len(*diff1))
		assert.Equal(t, 500, (*diff1)[0].Quantity.Int())
		assert.Equal(t, -200, (*diff1)[1].Quantity.Int())
		diff2 := ps2.Diff(ps1)
		assert.Equal(t, 2, len(*diff2))
		assert.Equal(t, 200, (*diff2)[0].Quantity.Int())
		assert.Equal(t, -500, (*diff2)[1].Quantity.Int())
		assert.Equal(t, 3, len(*ps1))
		assert.Equal(t, 3, len(*ps2))
		assert.Equal(t, 1200, ps1.sumQuantity())
		assert.Equal(t, 900, ps2.sumQuantity())
	})
}
func TestCopy(t *testing.T) {
	ps1 := NewPositions()
	p1, _ := NewPosition("spot_buy", "1234", 500)
	p2, _ := NewPosition("margin_buy", "1234", 400)
	*ps1 = append(*ps1, p1, p2)
	ps2 := ps1.copy()
	assert.Equal(t, 900, ps1.sumQuantity())
	assert.Equal(t, 900, ps2.sumQuantity())
	(*ps2)[0].Quantity -= 400
	assert.Equal(t, 500, ps2.sumQuantity())
	assert.Equal(t, 900, ps1.sumQuantity(), "コピー元に影響がない")
}
