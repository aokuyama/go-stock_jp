package ordertype

import (
	"errors"
)

type PayType string

func NewPayType(v string) (*PayType, error) {
	types := GetPayTypes()
	for _, t := range types {
		if v == t {
			p := PayType(v)
			return &p, nil
		}
	}
	return nil, errors.New("Invalid pay type:" + v)
}

func (p *PayType) String() string {
	return string(*p)
}

func (p *PayType) IsPosition() bool {
	return false
}

func (p *PayType) IsPay() bool {
	return true
}

func (p *PayType) IsSpot() bool {
	return p.String() == "spot_sell"
}

func (p *PayType) IsMargin() bool {
	return !p.IsSpot()
}

func GetPayTypes() [3]string {
	return [...]string{
		"spot_sell",
		"pay_sell",
		"pay_buy",
	}
}

func (p *PayType) IsEqual(t *PayType) bool {
	return p.String() == t.String()
}

func (p *PayType) PairPositionType() (*PositionType, error) {
	pt := getPairPositionTypes()
	return NewPositionType(pt[p.String()])
}

func getPairPositionTypes() map[string]string {
	return map[string]string{
		"spot_sell": "spot_buy",
		"pay_sell":  "margin_buy",
		"pay_buy":   "margin_sell",
	}
}
