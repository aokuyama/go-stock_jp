package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledOrderCondition(t *testing.T) {
	var o *OrderCondition
	var err error
	o, err = NewOrderCondition("normal")
	assert.Equal(t, "normal", o.String(), "有効な注文条件")
	assert.NoError(t, err)
	o, err = NewOrderCondition("open")
	assert.Equal(t, "open", o.String(), "有効な注文条件")
	assert.NoError(t, err)
	o, err = NewOrderCondition("close")
	assert.Equal(t, "close", o.String(), "有効な注文条件")
	assert.NoError(t, err)
	o, err = NewOrderCondition("close_bid")
	assert.Equal(t, "close_bid", o.String(), "有効な注文条件")
	assert.NoError(t, err)
}

func TestDisabledOrderCondition(t *testing.T) {
	var err error
	_, err = NewOrderCondition("standard")
	assert.Error(t, err)
	_, err = NewOrderCondition("1")
	assert.Error(t, err)
}
