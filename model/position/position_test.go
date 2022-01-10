package position

import (
	"testing"

	"github.com/aokuyama/go-stock_jp/model/trade"
	"github.com/stretchr/testify/assert"
)

func TestNewPosition(t *testing.T) {
	var p *Position
	var equal string
	p, _ = NewPosition("margin_buy", "1300", 100)
	equal = `{"position_type":"margin_buy","security_code":"1300","quantity":100}`
	assert.Equal(t, equal, string(p.String()), "有効なポジション")
	p, _ = NewPosition("spot_buy", "2349", 10000)
	equal = `{"position_type":"spot_buy","security_code":"2349","quantity":10000}`
	assert.Equal(t, equal, string(p.String()), "有効なポジション")
}
func TestErrorNewPosition(t *testing.T) {
	var err error
	_, err = NewPosition("margin_buy", "1300", -100)
	assert.Error(t, err)
	_, err = NewPosition("margin_buy", "300", 1000)
	assert.Error(t, err)
	_, err = NewPosition("pay_buy", "1300", 1000)
	assert.Error(t, err)
}

func TestIsEqualTarget(t *testing.T) {
	var p *Position
	p, _ = NewPosition("margin_buy", "1300", 100)
	t1, _ := trade.NewPositionTrade("margin_buy", "1300", 100)
	assert.True(t, p.IsEqualTarget(t1), "対象が同じ")
	t2, _ := trade.NewPositionTrade("margin_sell", "1300", 100)
	assert.True(t, p.IsEqualTarget(t2), "対象が同じなら関係ない取引でもtrue")
	t3, _ := trade.NewPayTrade("pay_buy", "1300", 100)
	assert.True(t, p.IsEqualTarget(t3), "対象が同じなら関係ない手仕舞い取引でもtrue")
}

func TestIsEqualPosition(t *testing.T) {
	var p *Position
	p, _ = NewPosition("margin_buy", "1300", 100)
	t1, _ := NewPosition("margin_buy", "1300", 100)
	assert.True(t, p.IsEqualPosition(t1), "対象が同じ")
	t2, _ := NewPosition("spot_buy", "1300", 100)
	assert.False(t, p.IsEqualPosition(t2), "対象が同じでも関係ない取引ならfalse")
}
func TestCompletePosition(t *testing.T) {
	var p *Position
	p, _ = NewPosition("margin_buy", "1300", 200)
	assert.False(t, p.IsCompleted())
	p.Quantity -= 100
	assert.False(t, p.IsCompleted())
	p.Quantity -= 100
	assert.True(t, p.IsCompleted())
	p.Quantity -= 100
	assert.False(t, p.IsCompleted())
}

func TestIntegratePosition(t *testing.T) {
	p1, _ := NewPosition("margin_buy", "1300", 200)
	p2, _ := NewPosition("margin_buy", "1300", 300)
	p3, _ := NewPosition("margin_buy", "1300", 100)
	p1.integrate(p2)
	assert.Equal(t, 500, p1.Quantity.Int())
	assert.Equal(t, 0, p2.Quantity.Int())
	p1.integrate(p2)
	assert.Equal(t, 500, p1.Quantity.Int())
	assert.Equal(t, 0, p2.Quantity.Int())
	p2.integrate(p3)
	assert.Equal(t, 500, p1.Quantity.Int())
	assert.Equal(t, 100, p2.Quantity.Int())
	assert.Equal(t, 0, p3.Quantity.Int())
	p1.integrate(p3)
	assert.Equal(t, 500, p1.Quantity.Int())
	assert.Equal(t, 100, p2.Quantity.Int())
	assert.Equal(t, 0, p3.Quantity.Int())
	p1.integrate(p2)
	assert.Equal(t, 600, p1.Quantity.Int())
	assert.Equal(t, 0, p2.Quantity.Int())
	assert.Equal(t, 0, p3.Quantity.Int())
}
func TestIntegrateError(t *testing.T) {
	var err error
	p1, _ := NewPosition("margin_buy", "1300", 200)
	p2, _ := NewPosition("spot_buy", "1300", 300)
	p3, _ := NewPosition("margin_sell", "1300", 100)
	p4, _ := NewPosition("margin_buy", "1301", 200)
	err = p1.integrate(p1)
	assert.Error(t, err, "自分自身が対象だと失敗")
	err = p1.integrate(p2)
	assert.Error(t, err, "タイプが違うと失敗")
	err = p1.integrate(p3)
	assert.Error(t, err, "タイプが違うと失敗")
	err = p1.integrate(p4)
	assert.Error(t, err, "ターゲットが違うと失敗")
}
func TestNewPositionByTrade(t *testing.T) {
	pt1, _ := trade.NewPositionTrade("margin_buy", "1300", 200)
	pos1 := NewPositionByTrade(pt1)
	assert.Equal(t, "{\"position_type\":\"margin_buy\",\"security_code\":\"1300\",\"quantity\":200}", pos1.String())
	pt2, _ := trade.NewPositionTrade("spot_buy", "5301", 300)
	pos2 := NewPositionByTrade(pt2)
	assert.Equal(t, "{\"position_type\":\"spot_buy\",\"security_code\":\"5301\",\"quantity\":300}", pos2.String())
}
func TestNewPositionByPayTrade(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		p1, _ := NewPosition("margin_buy", "1300", 200)
		pt1, _ := trade.NewPayTrade("pay_sell", "1300", 100)
		p2, _ := NewPositionByPayTrade(p1, pt1)
		assert.Equal(t, "{\"position_type\":\"margin_buy\",\"security_code\":\"1300\",\"quantity\":100}", p2.String())
		assert.Equal(t, 200, p1.Quantity.Int(), "元のポジションは変化していない")
		p3, _ := NewPositionByPayTrade(p2, pt1)
		assert.Equal(t, "{\"position_type\":\"margin_buy\",\"security_code\":\"1300\",\"quantity\":0}", p3.String())
		assert.Equal(t, 200, p1.Quantity.Int(), "元のポジションは変化していない")
	})
	t.Run("error", func(t *testing.T) {
		var err error
		pos, _ := NewPosition("margin_buy", "1301", 200)
		pt1, _ := trade.NewPayTrade("pay_buy", "1301", 100)
		_, err = NewPositionByPayTrade(pos, pt1)
		assert.Error(t, err, "タイプが違うと失敗")
		pt2, _ := trade.NewPayTrade("spot_sell", "1301", 100)
		_, err = NewPositionByPayTrade(pos, pt2)
		assert.Error(t, err, "タイプが違うと失敗")
		pt3, _ := trade.NewPayTrade("pay_sell", "1300", 100)
		_, err = NewPositionByPayTrade(pos, pt3)
		assert.Error(t, err, "ターゲットが違うと失敗")
		pt4, _ := trade.NewPayTrade("pay_sell", "1301", 300)
		_, err = NewPositionByPayTrade(pos, pt4)
		assert.Error(t, err, "数量が超過していると失敗")
	})
}
