package aggregate

import (
	"github.com/aokuyama/go-stock_jp/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewTrade(t *testing.T) {
	var ts Trades
	var tr entity.ITrade
	var equal string
	tr, _ = ts.AddNewTrade("margin_buy", "5651", 500)
	equal = `{"PositionType":"margin_buy","SecurityCode":"5651","Quantity":500}`
	assert.Equal(t, equal, string(tr.String()), "有効な取引")
	tr, _ = ts.AddNewTrade("spot_sell", "1392", 1200)
	equal = `{"PayType":"spot_sell","SecurityCode":"1392","Quantity":1200}`
	assert.Equal(t, equal, string(tr.String()), "有効な取引")
}
