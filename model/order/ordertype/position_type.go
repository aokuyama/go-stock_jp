package ordertype

import (
	"errors"
)

type PositionType string

func NewPositionType(v string) (*PositionType, error) {
	types := GetPositionTypes()
	for _, t := range types {
		if v == t {
			p := PositionType(v)
			return &p, nil
		}
	}
	return nil, errors.New("Invalid position type:" + v)
}

func (p *PositionType) String() string {
	return string(*p)
}

func (p *PositionType) IsPosition() bool {
	return true
}

func (p *PositionType) IsPay() bool {
	return false
}

func (p *PositionType) IsSpot() bool {
	return p.String() == "spot_buy"
}

func (p *PositionType) IsMargin() bool {
	return !p.IsSpot()
}

func GetPositionTypes() [3]string {
	return [...]string{
		"spot_buy",
		"margin_buy",
		"margin_sell",
	}
}

func (p *PositionType) IsEqual(t *PositionType) bool {
	return p.String() == t.String()
}

func (p *PositionType) IsCorrectPayType(t *PayType) bool {
	pt := getPairPayTypes()
	return pt[p.String()] == t.String()
}

func getPairPayTypes() map[string]string {
	return map[string]string{
		"spot_buy":    "spot_sell",
		"margin_buy":  "pay_sell",
		"margin_sell": "pay_buy",
	}
}
