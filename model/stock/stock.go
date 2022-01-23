package stock

import "encoding/json"

type Stock struct {
	SequrityCode SecurityCode
	Market       Market
}

func New(security_code string, market string) (*Stock, error) {
	s, err := NewSecurityCode(security_code)
	if err != nil {
		return nil, err
	}
	m, err := NewMarket(market)
	if err != nil {
		return nil, err
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
