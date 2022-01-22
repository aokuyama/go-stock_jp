package ordertype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdertype(t *testing.T) {
	var ot *Ordertype
	var err error
	ot, err = New("margin_buy", "general")
	assert.Equal(t, `{"Trade":"margin_buy","Margin":"general"}`, string(ot.String()), "有効なタイプ")
	assert.NoError(t, err)
	ot, err = New("spot_buy", "")
	assert.Equal(t, `{"Trade":"spot_buy","Margin":null}`, string(ot.String()), "有効なタイプ")
	assert.NoError(t, err)
}

func TestErrorOrdertype(t *testing.T) {
	var ot *Ordertype
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
