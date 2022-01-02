package value

import (
	"errors"
)

type Quantity int

func NewQuantity(v int) (*Quantity, error) {
	if v <= 0 {
		return nil, errors.New("Less than 0")
	}
	if v%100 != 0 {
		return nil, errors.New("Not divisible by 100")
	}
	q := Quantity(v)
	return &q, nil
}

func NewErrorQuantity(v int) (*Quantity, error) {
	q, err := NewQuantity(v)
	if err != nil {
		return nil, err
	}
	*q = 0
	*q -= Quantity(v)
	return q, nil
}

func (q *Quantity) IsError() bool {
	return *q < 0
}

func (q *Quantity) Int() int {
	return int(*q)
}
