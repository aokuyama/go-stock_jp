package order

import (
	"errors"
)

type MarginType string

func NewMarginType(v string) (*MarginType, error) {
	types := GetMarginTypes()
	for _, t := range types {
		if v == t {
			m := MarginType(v)
			return &m, nil
		}
	}
	return nil, errors.New("Invalid margin type:" + v)
}
func GetMarginTypes() [2]string {
	return [...]string{
		"system",
		"general",
	}
}
func (m *MarginType) String() string {
	return string(*m)
}
