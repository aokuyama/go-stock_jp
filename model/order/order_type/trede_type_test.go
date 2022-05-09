package order_type_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order/order_type"

	"github.com/stretchr/testify/assert"
)

func TestEnabledTradeType(t *testing.T) {
	var tt TradeType
	tt, _ = NewTradeType("spot_buy")
	assert.Equal(t, "spot_buy", tt.String(), "有効なタイプ")
	tt, _ = NewTradeType("margin_buy")
	assert.Equal(t, "margin_buy", tt.String(), "有効なタイプ")
	tt, _ = NewTradeType("pay_buy")
	assert.Equal(t, "pay_buy", tt.String(), "有効なタイプ")
}

func TestDisabledTradeType(t *testing.T) {
	var err error
	_, err = NewPayType("margin_sell")
	assert.Error(t, err)
	_, err = NewPayType("spot_selll")
	assert.Error(t, err)
}

func TestTradeTypeIsPositionOrPay(t *testing.T) {
	var tt TradeType
	tt, _ = NewTradeType("spot_buy")
	assert.True(t, tt.IsPosition())
	assert.False(t, tt.IsPay())
	assert.True(t, tt.IsSpot())
	assert.False(t, tt.IsMargin())
	tt, _ = NewTradeType("spot_sell")
	assert.True(t, tt.IsPay())
	assert.False(t, tt.IsPosition())
	assert.True(t, tt.IsSpot())
	assert.False(t, tt.IsMargin())
	tt, _ = NewTradeType("margin_buy")
	assert.True(t, tt.IsPosition())
	assert.False(t, tt.IsPay())
	assert.False(t, tt.IsSpot())
	assert.True(t, tt.IsMargin())
	tt, _ = NewTradeType("pay_sell")
	assert.True(t, tt.IsPay())
	assert.False(t, tt.IsPosition())
	assert.False(t, tt.IsSpot())
	assert.True(t, tt.IsMargin())
	tt, _ = NewTradeType("margin_sell")
	assert.True(t, tt.IsPosition())
	assert.False(t, tt.IsPay())
	assert.False(t, tt.IsSpot())
	assert.True(t, tt.IsMargin())
	tt, _ = NewTradeType("pay_buy")
	assert.True(t, tt.IsPay())
	assert.False(t, tt.IsPosition())
	assert.False(t, tt.IsSpot())
	assert.True(t, tt.IsMargin())
}
