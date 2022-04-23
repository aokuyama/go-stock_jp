package common

import (
	"errors"
)

type Locale string

func NewLocale(v string) (*Locale, error) {
	types := getLocales()
	for _, t := range types {
		if v == t {
			p := Locale(v)
			return &p, nil
		}
	}
	return nil, errors.New("invalid locale:" + v)
}

func (p *Locale) String() string {
	return string(*p)
}

func getLocales() [2]string {
	return [...]string{
		"jp",
		"ny",
	}
}
