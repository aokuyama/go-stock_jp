package value

import (
	"errors"
)

type OrderCondition string

func NewOrderCondition(v string) (*OrderCondition, error) {
	types := GetOrderConditions()
	for _, t := range types {
		if v == t {
			c := OrderCondition(v)
			return &c, nil
		}
	}
	return nil, errors.New("Invalid order condition:" + v)
}
func GetOrderConditions() [4]string {
	return [...]string{
		"normal",
		"open",
		"close",
		"close_bid",
	}
}
func (c *OrderCondition) String() string {
	return string(*c)
}
