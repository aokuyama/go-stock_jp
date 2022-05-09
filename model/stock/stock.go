package stock

import (
	"encoding/json"

	"github.com/aokuyama/go-generic_subdomains/errs"
)

type Stock struct {
	sequrityCode SecurityCode
	market       Market
}

type stockJson struct {
	SequrityCode SecurityCode `json:"security_code"`
	Market       Market       `json:"market"`
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

func (s *Stock) MarshalJSON() ([]byte, error) {
	return json.Marshal(&stockJson{
		SequrityCode: s.sequrityCode,
		Market:       s.market,
	})
}

func (s *Stock) UnmarshalJSON(b []byte) error {
	j := stockJson{}
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*s = Stock{
		sequrityCode: j.SequrityCode,
		market:       j.Market,
	}
	return nil
}
