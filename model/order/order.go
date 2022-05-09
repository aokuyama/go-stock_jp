package order

import (
	"encoding/json"

	"github.com/aokuyama/go-generic_subdomains/errs"
	"github.com/aokuyama/go-stock_jp/model/common"
	"github.com/aokuyama/go-stock_jp/model/order/order_type"
	"github.com/aokuyama/go-stock_jp/model/order/trigger"
	"github.com/aokuyama/go-stock_jp/model/stock"
)

type Order struct {
	ID        *OrderID             `json:"id"`
	Stock     stock.Stock          `json:"stock"`
	Type      order_type.OrderType `json:"type"`
	Condition Condition            `json:"condition"`
	Bid       *stock.StockPrice    `json:"bid"`
	Trigger   trigger.Trigger      `json:"trigger"`
	Quantity  Quantity             `json:"quantity"`
	Date      common.Date          `json:"date"`
	Session   Session              `json:"session"`
	Status    Status               `json:"status"`
	IsCancel  bool                 `json:"is_cancel"`
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
	ot, err := order_type.New(trade_type, margin_type)
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

func (o *Order) CanBeOrdered() bool {
	return o.Status == "not_ordered"
}

func (o *Order) Ordering() *Order {
	s, err := NewStatus("ordering")
	if err != nil {
		panic(err)
	}
	o2 := Order{
		o.ID,
		o.Stock,
		o.Type,
		o.Condition,
		o.Bid,
		o.Trigger,
		o.Quantity,
		o.Date,
		o.Session,
		*s,
		o.IsCancel,
	}
	return &o2
}
