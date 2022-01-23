package order

import (
	"encoding/json"

	"github.com/aokuyama/go-stock_jp/common"
	"github.com/aokuyama/go-stock_jp/model/order/ordertype"
	"github.com/aokuyama/go-stock_jp/model/order/trigger"
	"github.com/aokuyama/go-stock_jp/model/stock"
)

type Order struct {
	ID        *OrderID
	Stock     stock.Stock
	Type      ordertype.Ordertype
	Condition Condition
	Bid       *stock.StockPrice
	Trigger   trigger.Trigger
	Quantity  Quantity

	Date    common.Date
	Session Session

	Status   Status
	IsCancel bool
}

func New(
	id uint64,
	security_code string,
	market string,
	trade_type string,
	margin_type string,
	condition string,
	bid float64,
	trigger_type string,
	trigger_price float64,
	quantity int,
	date string,
	session string,
	status string,
	isCancel bool,
) (*Order, error) {
	var err error
	var i *OrderID
	var b *stock.StockPrice
	if id != 0 {
		i, err = NewOrderID(id)
		if err != nil {
			return nil, err
		}
	}
	s, err := stock.NewStock(security_code, market)
	if err != nil {
		return nil, err
	}
	ot, err := ordertype.New(trade_type, margin_type)
	if err != nil {
		return nil, err
	}
	c, err := NewCondition(condition)
	if err != nil {
		return nil, err
	}
	if bid != 0 {
		b, err = stock.NewStockPrice(bid)
		if err != nil {
			return nil, err
		}
	}
	t, err := trigger.New(trigger_type, trigger_price)
	if err != nil {
		return nil, err
	}
	q, err := NewQuantity(quantity)
	if err != nil {
		return nil, err
	}
	d, err := common.NewDate(date)
	if err != nil {
		return nil, err
	}
	se, err := NewSession(session)
	if err != nil {
		return nil, err
	}
	st, err := NewStatus(status)
	if err != nil {
		return nil, err
	}
	o := Order{
		i,
		*s,
		*ot,
		*c,
		b,
		*t,
		*q,
		*d,
		*se,
		*st,
		isCancel,
	}
	return &o, nil
}

func (o *Order) String() string {
	j, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}
