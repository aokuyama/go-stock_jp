package position_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/position"

	"github.com/aokuyama/go-stock_jp/model/trade"
	"github.com/stretchr/testify/assert"
)

func TestNewPosition(t *testing.T) {
	var p *Position
	var equal string
	p, _ = New("margin_buy", "1300", 100)
	equal = `{"position_type":"margin_buy","security_code":"1300","quantity":100}`
	assert.Equal(t, equal, string(p.String()), "有効なポジション")
	p, _ = New("spot_buy", "2349", 10000)
	equal = `{"position_type":"spot_buy","security_code":"2349","quantity":10000}`
	assert.Equal(t, equal, string(p.String()), "有効なポジション")
}
func TestErrorNewPosition(t *testing.T) {
	var err error
	_, err = New("margin_buy", "1300", -100)
	assert.Error(t, err)
	_, err = New("margin_buy", "300", 1000)
	assert.Error(t, err)
	_, err = New("pay_buy", "1300", 1000)
	assert.Error(t, err)
}

func TestIsEqualTarget(t *testing.T) {
	var p *Position
	p, _ = New("margin_buy", "1300", 100)
	t1, _ := trade.NewPositionTrade("margin_buy", "1300", 100)
	assert.True(t, p.IsEqualTarget(t1), "対象が同じ")
	t2, _ := trade.NewPositionTrade("margin_sell", "1300", 100)
	assert.True(t, p.IsEqualTarget(t2), "対象が同じなら関係ない取引でもtrue")
	t3, _ := trade.NewPayTrade("pay_buy", "1300", 100)
	assert.True(t, p.IsEqualTarget(t3), "対象が同じなら関係ない手仕舞い取引でもtrue")
}

func TestIsEqualPosition(t *testing.T) {
	var p *Position
	p, _ = New("margin_buy", "1300", 100)
	t1, _ := New("margin_buy", "1300", 100)
	assert.True(t, p.IsEqualPosition(t1), "対象が同じ")
	t2, _ := New("spot_buy", "1300", 100)
	assert.False(t, p.IsEqualPosition(t2), "対象が同じでも関係ない取引ならfalse")
}
func TestCompletePosition(t *testing.T) {
	var p *Position
	p, _ = New("margin_buy", "1300", 200)
	assert.False(t, p.IsCompleted())
	p2, _ := p.Decrement(100)
	assert.False(t, p2.IsCompleted())
	p3, _ := p2.Decrement(100)
	assert.True(t, p3.IsCompleted())
	_, err := p3.Decrement(100)
	assert.Error(t, err)
}
func TestNewPositionByTrade(t *testing.T) {
	pt1, _ := trade.NewPositionTrade("margin_buy", "1300", 200)
	pos1 := NewPositionByTrade(pt1)
	assert.Equal(t, "{\"position_type\":\"margin_buy\",\"security_code\":\"1300\",\"quantity\":200}", pos1.String())
	pt2, _ := trade.NewPositionTrade("spot_buy", "5301", 300)
	pos2 := NewPositionByTrade(pt2)
	assert.Equal(t, "{\"position_type\":\"spot_buy\",\"security_code\":\"5301\",\"quantity\":300}", pos2.String())
}
func TestNewPositionByPositionAndPayTrade(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		p1, _ := New("margin_buy", "1300", 200)
		pt1, _ := trade.NewPayTrade("pay_sell", "1300", 100)
		p2, _ := NewPositionByPositionAndPayTrade(p1, pt1)
		assert.Equal(t, "{\"position_type\":\"margin_buy\",\"security_code\":\"1300\",\"quantity\":100}", p2.String())
		assert.Equal(t, 200, p1.Quantity(), "元のポジションは変化していない")
		p3, _ := NewPositionByPositionAndPayTrade(p2, pt1)
		assert.Equal(t, "{\"position_type\":\"margin_buy\",\"security_code\":\"1300\",\"quantity\":0}", p3.String())
		assert.Equal(t, 200, p1.Quantity(), "元のポジションは変化していない")
	})
	t.Run("error", func(t *testing.T) {
		var err error
		pos, _ := New("margin_buy", "1301", 200)
		pt1, _ := trade.NewPayTrade("pay_buy", "1301", 100)
		_, err = NewPositionByPositionAndPayTrade(pos, pt1)
		assert.Error(t, err, "タイプが違うと失敗")
		pt2, _ := trade.NewPayTrade("spot_sell", "1301", 100)
		_, err = NewPositionByPositionAndPayTrade(pos, pt2)
		assert.Error(t, err, "タイプが違うと失敗")
		pt3, _ := trade.NewPayTrade("pay_sell", "1300", 100)
		_, err = NewPositionByPositionAndPayTrade(pos, pt3)
		assert.Error(t, err, "ターゲットが違うと失敗")
		pt4, _ := trade.NewPayTrade("pay_sell", "1301", 300)
		_, err = NewPositionByPositionAndPayTrade(pos, pt4)
		assert.Error(t, err, "数量が超過していると失敗")
	})
}
func TestNewIntegratePosition(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		p1, _ := New("margin_buy", "1300", 200)
		p2, _ := New("margin_buy", "1300", 100)
		p3, _ := New("margin_buy", "1300", 500)
		p12, _ := NewIntegratePosition(p1, p2)
		assert.Equal(t, 300, p12.Quantity(), "統合されたポジション")
		p23, _ := NewIntegratePosition(p2, p3)
		assert.Equal(t, 600, p23.Quantity(), "統合されたポジション")
		assert.Equal(t, 200, p1.Quantity(), "元のポジションは変化していない")
		assert.Equal(t, 100, p2.Quantity(), "元のポジションは変化していない")
		assert.Equal(t, 500, p3.Quantity(), "元のポジションは変化していない")
	})
	t.Run("error", func(t *testing.T) {
		var err error
		pos, _ := New("margin_buy", "1301", 200)
		pt1, _ := New("margin_sell", "1301", 100)
		_, err = NewIntegratePosition(pos, pt1)
		assert.Error(t, err, "タイプが違うと失敗")
		pt2, _ := New("spot_buy", "1301", 100)
		_, err = NewIntegratePosition(pos, pt2)
		assert.Error(t, err, "タイプが違うと失敗")
		pt3, _ := New("margin_buy", "1300", 100)
		_, err = NewIntegratePosition(pos, pt3)
		assert.Error(t, err, "ターゲットが違うと失敗")
	})
}
