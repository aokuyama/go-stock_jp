package stock

import (
	"errors"
)

type StockPrice float64

func NewStockPrice(v float64) (*StockPrice, error) {
	if v <= 0 {
		return nil, errors.New("0 or less")
	}
	if v != float64(float64(int(v*10))/10) {
		return nil, errors.New("2 or more decimal places")
	}
	p := StockPrice(v)
	return &p, nil
}

func (p *StockPrice) Float() float64 {
	return float64(*p)
}
