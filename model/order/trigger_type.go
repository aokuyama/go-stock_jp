package order

import (
	"errors"
)

type TriggerType string

func NewTriggerType(v string) (*TriggerType, error) {
	types := GetTriggerTypes()
	for _, t := range types {
		if v == t {
			tt := TriggerType(v)
			return &tt, nil
		}
	}
	return nil, errors.New("Invalid trigger type:" + v)
}
func GetTriggerTypes() [2]string {
	return [...]string{
		"more",
		"less",
	}
}
func (tt *TriggerType) String() string {
	return string(*tt)
}
