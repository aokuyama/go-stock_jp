package stock

import (
	"errors"
	"unicode/utf8"
)

type SecurityCode string

func NewSecurityCode(v string) (*SecurityCode, error) {
	if utf8.RuneCountInString(v) != 4 {
		return nil, errors.New("証券コードは4文字")
	}
	c := SecurityCode(v)
	return &c, nil
}

func (p *SecurityCode) String() string {
	return string(*p)
}

func (s *SecurityCode) IsEqual(security_code *SecurityCode) bool {
	return s.String() == security_code.String()
}
