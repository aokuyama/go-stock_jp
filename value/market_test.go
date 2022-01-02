package value

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
