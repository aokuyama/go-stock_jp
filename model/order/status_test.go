package order_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order"

	"github.com/stretchr/testify/assert"
)

func TestEnabledStatus(t *testing.T) {
	var st *Status
	var err error
	st, err = NewStatus("not_ordered")
	assert.NoError(t, err)
	assert.Equal(t, "not_ordered", st.String(), "有効なステータス")
	st, err = NewStatus("completed")
	assert.NoError(t, err)
	assert.Equal(t, "completed", st.String(), "有効なステータス")
	st, err = NewStatus("ordered")
	assert.NoError(t, err)
	assert.Equal(t, "ordered", st.String(), "有効なステータス")
}

func TestDisabledStatus(t *testing.T) {
	var st *Status
	var err error
	_, err = NewStatus("margin_sell")
	assert.Nil(t, st)
	assert.Error(t, err)
	_, err = NewStatus("spot_selll")
	assert.Nil(t, st)
	assert.Error(t, err)
}
