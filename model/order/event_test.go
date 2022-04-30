package order_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order"
	"github.com/aokuyama/go-stock_jp/test"
	"github.com/stretchr/testify/assert"
)

func TestOrderingEvent(t *testing.T) {
	o := test.GetOrder()
	e := NewOrderingEvent(o)
	assert.Equal(t, `{"id":null,"stock":{"security_code":"3662","market":"jpx"},"type":{"trade":"margin_buy","margin":"system"},"condition":"normal","bid":null,"trigger":{"type":"more","price":10000},"quantity":200,"date":"2022-01-24","session":"afternoon","status":"not_ordered","is_cancel":false}`, e.Data())
}
