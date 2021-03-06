package order

import (
	"errors"
)

type Quantity int

func NewQuantity(v int) (*Quantity, error) {
	if v <= 0 {
		return nil, errors.New("0 or less")
	}
	if v%100 != 0 {
		return nil, errors.New("not divisible by 100")
	}
	q := Quantity(v)
	return &q, nil
}

func (q *Quantity) Int() int {
	return int(*q)
}
