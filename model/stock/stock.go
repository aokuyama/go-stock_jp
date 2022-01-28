package stock

import (
	"encoding/json"

	"github.com/aokuyama/go-generic_subdomains/errs"
)

type Stock struct {
	SequrityCode SecurityCode
	Market       Market
}

func New(security_code string, market string) (*Stock, error) {
	errs := errs.New()
	s, err := NewSecurityCode(security_code)
	errs.Append(err)
	m, err := NewMarket(market)
	errs.Append(err)
	if errs.Err() != nil {
		return nil, errs.Err()
	}
	return &Stock{
		*s,
		*m,
	}, nil
}

func (s *Stock) String() string {
	j, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return (string)(j)
}
