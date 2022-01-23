package order_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order"
	"github.com/stretchr/testify/assert"
)

func TestOrderID(t *testing.T) {
	var i *OrderID
	var err error
	i, err = NewOrderID(1)
	assert.Equal(t, "1", i.String(), "有効なID")
	assert.NoError(t, err)
	i, err = NewOrderID(1234567890)
	assert.Equal(t, "1234567890", i.String(), "有効なID")
	assert.NoError(t, err)
}

func TestErrorOrderID(t *testing.T) {
	var i *OrderID
	var err error
	i, err = NewOrderID(0)
	assert.Nil(t, i)
	assert.Error(t, err)
}
