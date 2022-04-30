package trigger

import (
	"encoding/json"
	"errors"

	"github.com/aokuyama/go-generic_subdomains/errs"
	"github.com/aokuyama/go-stock_jp/model/stock"
)

type Trigger struct {
	Type  *TriggerType      `json:"type"`
	Price *stock.StockPrice `json:"price"`
}

func New(ttype string, price float64) (*Trigger, error) {
	var t *TriggerType
	var p *stock.StockPrice
	var err error
	errs := errs.New()
	if ttype != "" {
		t, err = NewTriggerType(ttype)
		errs.Append(err)
	}
	if price != 0 {
		p, err = stock.NewStockPrice(price)
		errs.Append(err)
	}
	if errs.Err() != nil {
		return nil, errs.Err()
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
