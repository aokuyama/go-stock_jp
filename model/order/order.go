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
	id         *OrderID
	stock      stock.Stock
	order_type order_type.OrderType
	condition  Condition
	bid        *stock.StockPrice
	trigger    trigger.Trigger
	quantity   Quantity
	date       common.Date
	session    Session
	status     Status
	isCancel   bool
}
type orderJson struct {
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

func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(&orderJson{
		ID:        o.id,
		Stock:     o.stock,
		Type:      o.order_type,
		Condition: o.condition,
		Bid:       o.bid,
		Trigger:   o.trigger,
		Quantity:  o.quantity,
		Date:      o.date,
		Session:   o.session,
		Status:    o.status,
		IsCancel:  o.isCancel,
	})
}

func (o *Order) UnmarshalJSON(b []byte) error {
	j := orderJson{}
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*o = Order{
		id:         j.ID,
		stock:      j.Stock,
		order_type: j.Type,
		condition:  j.Condition,
		bid:        j.Bid,
		trigger:    j.Trigger,
		quantity:   j.Quantity,
		date:       j.Date,
		session:    j.Session,
		status:     j.Status,
		isCancel:   j.IsCancel,
	}
	return nil
}

func (o *Order) CanBeOrdered() bool {
	return o.status == "not_ordered"
}

func (o *Order) Ordering() *Order {
	s, err := NewStatus("ordering")
	if err != nil {
		panic(err)
	}
	o2 := Order{
		o.id,
		o.stock,
		o.order_type,
		o.condition,
		o.bid,
		o.trigger,
		o.quantity,
		o.date,
		o.session,
		*s,
		o.isCancel,
	}
	return &o2
}

func (o *Order) Status() string {
	return o.status.String()
}
