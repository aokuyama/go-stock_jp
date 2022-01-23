package trigger

import (
	"encoding/json"
	"errors"

	"github.com/aokuyama/go-stock_jp/model/stock"
)

type Trigger struct {
	Type  *TriggerType
	Price *stock.StockPrice
}

func New(ttype string, price float64) (*Trigger, error) {
	var t *TriggerType
	var p *stock.StockPrice
	var err error
	t, err = NewTriggerType(ttype)
	if ttype != "" {
		if err != nil {
			return nil, err
		}
	}
	p, err = stock.NewStockPrice(price)
	if price != 0 {
		if err != nil {
			return nil, err
		}
	}
	trigger := Trigger{t, p}
	if trigger.Error() != "" {
		return nil, errors.New(trigger.Error())
	}
	return &trigger, nil
}

func (t *Trigger) Error() string {
	if t.Type != nil {
		if t.Price == nil {
			return "undefined price"
		}
	}
	if t.Price != nil {
		if t.Type == nil {
			return "undefined type"
		}
	}
	return ""
}

func (t *Trigger) String() string {
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}
