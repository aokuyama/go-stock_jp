package order

import (
	"errors"
)

type Condition string

func NewCondition(v string) (*Condition, error) {
	types := GetConditions()
	for _, t := range types {
		if v == t {
			c := Condition(v)
			return &c, nil
		}
	}
	return nil, errors.New("Invalid order condition:" + v)
}
func GetConditions() [4]string {
	return [...]string{
		"normal",
		"open",
		"close",
		"close_bid",
	}
}
func (c *Condition) String() string {
	return string(*c)
}
