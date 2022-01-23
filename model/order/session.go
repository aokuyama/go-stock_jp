package order

import (
	"errors"
)

type Session string

func NewSession(v string) (*Session, error) {
	types := getSessions()
	for _, t := range types {
		if v == t {
			p := Session(v)
			return &p, nil
		}
	}
	return nil, errors.New("invalid session:" + v)
}

func (p *Session) String() string {
	return string(*p)
}

func getSessions() [2]string {
	return [...]string{
		"morning",
		"afternoon",
	}
}
