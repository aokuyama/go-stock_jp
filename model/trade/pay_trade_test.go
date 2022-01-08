package trade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPayTrade(t *testing.T) {
	var p *PayTrade
	var equal string
	p, _ = NewPayTrade("pay_sell", "5651", 500)
	equal = `{"PayType":"pay_sell","SecurityCode":"5651","Quantity":500}`
	assert.Equal(t, equal, string(p.String()), "有効な取引")
	p, _ = NewPayTrade("pay_buy", "1392", 1200)
	equal = `{"PayType":"pay_buy","SecurityCode":"1392","Quantity":1200}`
	assert.Equal(t, equal, string(p.String()), "有効な取引")
}
