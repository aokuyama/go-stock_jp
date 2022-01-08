package trade

import (
	"github.com/aokuyama/go-stock_jp/model/order"
	"github.com/aokuyama/go-stock_jp/model/stock"
)

type ITrade interface {
	String() string
	Target() *stock.SecurityCode
}

type IPosition interface {
	String() string
	Target() *stock.SecurityCode
	Type() *order.PositionType
}
