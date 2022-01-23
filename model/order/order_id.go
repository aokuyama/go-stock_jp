package order

import (
	"errors"
	"strconv"
)

type OrderID uint64

func NewOrderID(v uint64) (*OrderID, error) {
	if v <= 0 {
		return nil, errors.New("")
	}
	i := OrderID(v)
	return &i, nil
}

func (i *OrderID) String() string {
	return strconv.FormatUint(uint64(*i), 10)
}
