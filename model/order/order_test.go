package order_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order"
	"github.com/stretchr/testify/assert"
)

func TestEnableOrder(t *testing.T) {
	var o *Order
	var err error
	o, err = New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "completed", false)
	assert.Equal(t, `{"id":null,"stock":{"security_code":"9856","market":"jpx"},"type":{"trade":"margin_buy","margin":"system"},"condition":"normal","bid":null,"trigger":{"type":"more","price":100},"quantity":200,"date":"2022-01-24","session":"afternoon","status":"completed","is_cancel":false}`, string(o.String()), "有効なオーダー")
	assert.NoError(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", 100, "", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Equal(t, `{"id":1,"stock":{"security_code":"1324","market":"jpx"},"type":{"trade":"spot_buy","margin":null},"condition":"open","bid":100,"trigger":{"type":null,"price":null},"quantity":100,"date":"2022-01-23","session":"morning","status":"not_ordered","is_cancel":true}`, string(o.String()), "有効なオーダー")
	assert.NoError(t, err)
}

func TestErrorOrder(t *testing.T) {
	var o *Order
	var err error
	o, err = New(1, "132", "jpx", "spot_buy", "", "open", 5, "", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "", "spot_buy", "", "open", 5, "", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "", "", "open", 5, "", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "system", "open", 5, "", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "", 5, "", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", -1, "", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", 5, "a", 0, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", 5, "", 1, 100, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", 5, "", 0, 0, "2022-01-23", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", 5, "", 0, 100, "", "morning", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", 5, "", 0, 100, "2022-01-23", "", "not_ordered", true)
	assert.Nil(t, o)
	assert.Error(t, err)
	o, err = New(1, "1324", "jpx", "spot_buy", "", "open", 5, "", 0, 100, "2022-01-23", "", "", true)
	assert.Nil(t, o)
	assert.Error(t, err)
}
func TestCanBeOrdered(t *testing.T) {
	var o *Order
	o, _ = New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "not_ordered", false)
	assert.True(t, o.CanBeOrdered())
	o, _ = New(1, "1324", "jpx", "spot_buy", "", "open", 100, "", 0, 100, "2022-01-23", "morning", "completed", true)
	assert.False(t, o.CanBeOrdered())
}
