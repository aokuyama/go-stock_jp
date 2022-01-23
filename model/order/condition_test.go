package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledCondition(t *testing.T) {
	var o *Condition
	var err error
	o, err = NewCondition("normal")
	assert.Equal(t, "normal", o.String(), "有効な注文条件")
	assert.NoError(t, err)
	o, err = NewCondition("open")
	assert.Equal(t, "open", o.String(), "有効な注文条件")
	assert.NoError(t, err)
	o, err = NewCondition("close")
	assert.Equal(t, "close", o.String(), "有効な注文条件")
	assert.NoError(t, err)
	o, err = NewCondition("close_bid")
	assert.Equal(t, "close_bid", o.String(), "有効な注文条件")
	assert.NoError(t, err)
}

func TestDisabledCondition(t *testing.T) {
	var err error
	_, err = NewCondition("standard")
	assert.Error(t, err)
	_, err = NewCondition("1")
	assert.Error(t, err)
}
