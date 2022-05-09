package order_type_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order/order_type"

	"github.com/stretchr/testify/assert"
)

func TestOrderType(t *testing.T) {
	var ot *OrderType
	var err error
	ot, err = New("margin_buy", "general")
	assert.Equal(t, `{"trade":"margin_buy","margin":"general"}`, string(ot.String()), "有効なタイプ")
	assert.NoError(t, err)
	ot, err = New("spot_buy", "")
	assert.Equal(t, `{"trade":"spot_buy","margin":null}`, string(ot.String()), "有効なタイプ")
	assert.NoError(t, err)
}

func TestErrorOrderType(t *testing.T) {
	var ot *OrderType
	var err error
	ot, err = New("margin_buy", "a")
	assert.Nil(t, ot)
	assert.Error(t, err)
	ot, err = New("aaa", "system")
	assert.Nil(t, ot)
	assert.Error(t, err)

	ot, err = New("margin_sell", "")
	assert.Nil(t, ot)
	assert.Error(t, err, "信用タイプ指定がない")
	ot, err = New("pay_sell", "")
	assert.Nil(t, ot)
	assert.Error(t, err, "信用タイプ指定がない")
	ot, err = New("spot_sell", "general")
	assert.Nil(t, ot)
	assert.Error(t, err, "不要な信用タイプ指定がある")
}
