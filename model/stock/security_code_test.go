package stock_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/stock"

	"github.com/stretchr/testify/assert"
)

func TestEnabledCode(t *testing.T) {
	c, _ := NewSecurityCode("1300")
	assert.Equal(t, "1300", c.String(), "有効なセキュリティコード")
	c2, _ := NewSecurityCode("9999")
	assert.Equal(t, "9999", c2.String(), "有効なセキュリティコード")
}

func TestDisabledCode(t *testing.T) {
	var err error
	_, err = NewSecurityCode("900")
	assert.Error(t, err)
	_, err = NewSecurityCode("10000")
	assert.Error(t, err)
	_, err = NewSecurityCode("")
	assert.Error(t, err)
}

func TestEqualCode(t *testing.T) {
	c, _ := NewSecurityCode("9000")
	c2, _ := NewSecurityCode("9000")
	assert.True(t, c.IsEqual(c))
	assert.True(t, c.IsEqual(c2))
	c3, _ := NewSecurityCode("9001")
	assert.False(t, c.IsEqual(c3))
	ss := SecurityCode("")
	assert.False(t, c.IsEqual((&ss)))
}
