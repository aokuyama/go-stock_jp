package order_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order"
	"github.com/stretchr/testify/assert"
)

func TestPickupOrderingOrders(t *testing.T) {
	var orders *Collection
	o1, _ := New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "not_ordered", true)
	o2, _ := New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "stopped", true)
	o3, _ := New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "ordering", true)
	o4, _ := New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "ordered", true)
	o5, _ := New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "canceling", true)
	o6, _ := New(0, "9856", "jpx", "margin_buy", "system", "normal", 0, "more", 100, 200, "2022-01-24", "afternoon", "completed", true)
	orders, _ = NewCollection(o1, o2, o3, o4, o5, o6)
	assert.Equal(t, 6, len(*orders))
	assert.Equal(t, 1, len(*orders.PickupOrderingOrders()))
	orders, _ = NewCollection(o1, o1, o3, o4, o5)
	assert.Equal(t, 5, len(*orders))
	assert.Equal(t, 2, len(*orders.PickupOrderingOrders()))
}
