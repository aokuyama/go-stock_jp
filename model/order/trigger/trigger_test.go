package trigger_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order/trigger"

	"github.com/stretchr/testify/assert"
)

func TestEnableTrigger(t *testing.T) {
	var tt *Trigger
	var err error
	tt, err = New("more", 1.1)
	assert.Equal(t, `{"Type":"more","Price":1.1}`, string(tt.String()), "有効な逆指値")
	assert.NoError(t, err)
	tt, err = New("less", 10000)
	assert.Equal(t, `{"Type":"less","Price":10000}`, string(tt.String()), "有効な逆指値")
	assert.NoError(t, err)
	tt, err = New("", 0)
	assert.Equal(t, `{"Type":null,"Price":null}`, string(tt.String()), "逆指値なし")
	assert.NoError(t, err)
}

func TestDisableTrigger(t *testing.T) {
	var tt *Trigger
	var err error
	tt, err = New("less", -1)
	assert.Nil(t, tt)
	assert.Error(t, err)
	tt, err = New("aaa", 100)
	assert.Nil(t, tt)
	assert.Error(t, err)

	tt, err = New("more", 0)
	assert.Nil(t, tt)
	assert.Error(t, err, "値段指定がない")
	tt, err = New("", 5)
	assert.Nil(t, tt)
	assert.Error(t, err, "タイプ指定がない")
}
