package trade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPositionTrade(t *testing.T) {
	var p *PositionTrade
	var equal string
	p, _ = NewPositionTrade("spot_buy", "5651", 500)
	equal = `{"PositionType":"spot_buy","SecurityCode":"5651","Quantity":500}`
	assert.Equal(t, equal, string(p.String()), "有効な取引")
	p, _ = NewPositionTrade("margin_sell", "1392", 1200)
	equal = `{"PositionType":"margin_sell","SecurityCode":"1392","Quantity":1200}`
	assert.Equal(t, equal, string(p.String()), "有効な取引")
}
