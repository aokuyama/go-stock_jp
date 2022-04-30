package order

import (
	"errors"
)

type Status string

func NewStatus(v string) (*Status, error) {
	types := getStatuses()
	for _, t := range types {
		if v == t {
			p := Status(v)
			return &p, nil
		}
	}
	return nil, errors.New("invalid status:" + v)
}

func (p *Status) String() string {
	return string(*p)
}

func getStatuses() [6]string {
	return [...]string{
		"not_ordered",
		"stopped",
		"ordering",
		"ordered",
		"canceling",
		"completed",
	}
}
