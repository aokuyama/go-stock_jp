package order

import (
	"encoding/json"

	"github.com/aokuyama/go-generic_subdomains/errs"
	"github.com/aokuyama/go-stock_jp/model/common"
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
	var errs = errs.New()
	var i *OrderID
	var b *stock.StockPrice
	if id != 0 {
		i, err = NewOrderID(id)
		errs.Append(err)
	}
	s, err := stock.New(security_code, market)
	errs.Append(err)
	ot, err := ordertype.New(trade_type, margin_type)
	errs.Append(err)
	c, err := NewCondition(condition)
	errs.Append(err)
	if bid != 0 {
		b, err = stock.NewStockPrice(bid)
		errs.Append(err)
	}
	t, err := trigger.New(trigger_type, trigger_price)
	errs.Append(err)
	q, err := NewQuantity(quantity)
	errs.Append(err)
	d, err := common.NewDate(date)
	errs.Append(err)
	se, err := NewSession(session)
	errs.Append(err)
	st, err := NewStatus(status)
	errs.Append(err)
	if errs.Err() != nil {
		return nil, errs.Err()
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
