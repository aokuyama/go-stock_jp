package entity

import (
	"testing"

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
	_, err = NewPosition("margin_buy", "1300", 0)
	assert.Error(t, err)
	_, err = NewPosition("margin_buy", "300", 1000)
	assert.Error(t, err)
	_, err = NewPosition("pay_buy", "1300", 1000)
	assert.Error(t, err)
}

func TestIsEqualTarget(t *testing.T) {
	var p *Position
	p, _ = NewPosition("margin_buy", "1300", 100)
	t1, _ := NewPositionTrade("margin_buy", "1300", 100)
	assert.True(t, p.IsEqualTarget(t1), "対象が同じ")
	t2, _ := NewPositionTrade("margin_sell", "1300", 100)
	assert.True(t, p.IsEqualTarget(t2), "対象が同じなら関係ない取引でもtrue")
	t3, _ := NewPayTrade("pay_buy", "1300", 100)
	assert.True(t, p.IsEqualTarget(t3), "対象が同じなら関係ない手仕舞い取引でもtrue")
}

func TestIsEqualPosition(t *testing.T) {
	var p *Position
	p, _ = NewPosition("margin_buy", "1300", 100)
	t1, _ := NewPositionTrade("margin_buy", "1300", 100)
	assert.True(t, p.IsEqualPosition(t1), "対象が同じ")
	t2, _ := NewPositionTrade("spot_buy", "1300", 100)
	assert.False(t, p.IsEqualPosition(t2), "対象が同じでも関係ない取引ならfalse")
}

func TestIncludePosition(t *testing.T) {
	var err error
	p, _ := NewPosition("margin_buy", "1300", 400)
	assert.Equal(t, 400, p.Quantity.Int())
	t1, _ := NewPositionTrade("margin_buy", "1300", 300)
	p.IncludePosition(t1)
	assert.Equal(t, 700, p.Quantity.Int(), "ポジション量が増える")
	p.IncludePosition(t1)
	assert.Equal(t, 1000, p.Quantity.Int(), "ポジション量が増える")
	t2, _ := NewPositionTrade("margin_buy", "1301", 100)
	err = p.IncludePosition(t2)
	assert.Equal(t, 1000, p.Quantity.Int())
	assert.Error(t, err, "関係ない取引は無視")
	t3, _ := NewPositionTrade("margin_sell", "1300", 100)
	err = p.IncludePosition(t3)
	assert.Equal(t, 1000, p.Quantity.Int())
	assert.Error(t, err, "関係ない取引は無視")
}

func TestNewErrorPosition(t *testing.T) {
	var p *PayTrade
	var ep *Position
	p, _ = NewPayTrade("spot_sell", "3265", 100)
	ep, _ = NewErrorPosition(p)
	assert.True(t, ep.IsEqualPay(p))
	assert.Equal(t, -100, ep.Quantity.Int())
	assert.True(t, ep.IsError())
	p, _ = NewPayTrade("pay_buy", "4265", 1200)
	ep, _ = NewErrorPosition(p)
	assert.True(t, ep.IsEqualPay(p))
	assert.Equal(t, -1200, ep.Quantity.Int())
	assert.True(t, ep.IsError())
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
func TestOffset(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		p1, _ := NewPosition("margin_buy", "1300", 200)
		p2, _ := NewPosition("margin_buy", "1300", 300)
		p3, _ := NewPosition("margin_buy", "1300", 100)
		p2.offset(p1)
		assert.Equal(t, 0, p1.Quantity.Int())
		assert.Equal(t, 100, p2.Quantity.Int())
		p2.offset(p3)
		assert.Equal(t, 0, p1.Quantity.Int())
		assert.Equal(t, 0, p2.Quantity.Int())
		assert.Equal(t, 0, p3.Quantity.Int())
	})
	t.Run("error", func(t *testing.T) {
		var err error
		p1, _ := NewPosition("margin_buy", "1300", 200)
		p2, _ := NewPosition("spot_buy", "1300", 300)
		p3, _ := NewPosition("margin_sell", "1300", 100)
		p4, _ := NewPosition("margin_buy", "1301", 200)
		err = p1.offset(p1)
		assert.Error(t, err, "自分自身が対象だと失敗")
		err = p1.offset(p2)
		assert.Error(t, err, "タイプが違うと失敗")
		err = p1.offset(p3)
		assert.Error(t, err, "タイプが違うと失敗")
		err = p1.offset(p4)
		assert.Error(t, err, "ターゲットが違うと失敗")
	})
}
