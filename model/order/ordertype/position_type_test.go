package ordertype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledPositionType(t *testing.T) {
	var p *PositionType
	p, _ = NewPositionType("spot_buy")
	assert.Equal(t, "spot_buy", p.String(), "有効なポジションタイプ")
	p, _ = NewPositionType("margin_buy")
	assert.Equal(t, "margin_buy", p.String(), "有効なポジションタイプ")
	p, _ = NewPositionType("margin_sell")
	assert.Equal(t, "margin_sell", p.String(), "有効なポジションタイプ")
}

func TestDisabledPositionType(t *testing.T) {
	var err error
	_, err = NewPositionType("margin_buyy")
	assert.Error(t, err)
	_, err = NewPositionType("pay_buy")
	assert.Error(t, err)
}

func TestEqualPosition(t *testing.T) {
	p1, _ := NewPositionType("spot_buy")
	p2, _ := NewPositionType("spot_buy")
	assert.True(t, p1.IsEqual(p1))
	assert.True(t, p1.IsEqual(p2))
	p3, _ := NewPositionType("margin_sell")
	assert.False(t, p1.IsEqual(p3))
	p4 := PositionType("")
	assert.False(t, p1.IsEqual(&p4))
}

func TestIsCorrectPayType(t *testing.T) {
	sb, _ := NewPositionType("spot_buy")
	mb, _ := NewPositionType("margin_buy")
	ms, _ := NewPositionType("margin_sell")
	ss, _ := NewPayType("spot_sell")
	ps, _ := NewPayType("pay_sell")
	pb, _ := NewPayType("pay_buy")

	assert.True(t, sb.IsCorrectPayType(ss))
	assert.False(t, sb.IsCorrectPayType(ps))
	assert.False(t, sb.IsCorrectPayType(pb))

	assert.False(t, mb.IsCorrectPayType(ss))
	assert.True(t, mb.IsCorrectPayType(ps))
	assert.False(t, mb.IsCorrectPayType(pb))

	assert.False(t, ms.IsCorrectPayType(ss))
	assert.False(t, ms.IsCorrectPayType(ps))
	assert.True(t, ms.IsCorrectPayType(pb))
}
