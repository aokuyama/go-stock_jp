package trade

import (
	"github.com/aokuyama/go-stock_jp/model/stock"
)

type ITrade interface {
	String() string
	Target() *stock.SecurityCode
}
