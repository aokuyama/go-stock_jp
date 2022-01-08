package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledPayType(t *testing.T) {
	var p *PayType
	p, _ = NewPayType("spot_sell")
	assert.Equal(t, "spot_sell", p.String(), "有効な手仕舞いタイプ")
	p, _ = NewPayType("pay_buy")
	assert.Equal(t, "pay_buy", p.String(), "有効な手仕舞いタイプ")
	p, _ = NewPayType("pay_sell")
	assert.Equal(t, "pay_sell", p.String(), "有効な手仕舞いタイプ")
}

func TestDisabledPayType(t *testing.T) {
	var err error
	_, err = NewPayType("margin_sell")
	assert.Error(t, err)
	_, err = NewPayType("spot_selll")
	assert.Error(t, err)
}

func TestEqualPay(t *testing.T) {
	p1, _ := NewPayType("spot_sell")
	p2, _ := NewPayType("spot_sell")
	assert.True(t, p1.IsEqual(p1))
	assert.True(t, p1.IsEqual(p2))
	p3, _ := NewPayType("pay_buy")
	assert.False(t, p1.IsEqual(p3))
	p4 := PayType("")
	assert.False(t, p1.IsEqual(&p4))
}

func TestGetPositionType(t *testing.T) {
	var py *PayType
	var po *PositionType
	py, _ = NewPayType("spot_sell")
	po, _ = py.PairPositionType()
	assert.Equal(t, "spot_buy", po.String())
	py, _ = NewPayType("pay_sell")
	po, _ = py.PairPositionType()
	assert.Equal(t, "margin_buy", po.String())
	py, _ = NewPayType("pay_buy")
	po, _ = py.PairPositionType()
	assert.Equal(t, "margin_sell", po.String())
	pe := PayType("")
	_, err := pe.PairPositionType()
	assert.Error(t, err)
}
