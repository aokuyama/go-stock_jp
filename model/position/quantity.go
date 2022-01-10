package position

import (
	"errors"
)

type Quantity int

func NewQuantity(v int) (*Quantity, error) {
	if v < 0 {
		return nil, errors.New("less than 0")
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

func (q *Quantity) IsEqual(quantity *Quantity) bool {
	return q.Int() == quantity.Int()
}
