package stock_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/stock"
	"github.com/stretchr/testify/assert"
)

func TestEnabledMarket(t *testing.T) {
	var m *Market
	m, _ = NewMarket("jpx")
	assert.Equal(t, "jpx", string(*m), "有効な市場")
	m, _ = NewMarket("jasdaq")
	assert.Equal(t, "jasdaq", string(*m), "有効な市場")
}

func TestDisabledMarket(t *testing.T) {
	var err error
	_, err = NewMarket("abc")
	assert.Error(t, err)
	_, err = NewMarket("12345")
	assert.Error(t, err)
}
