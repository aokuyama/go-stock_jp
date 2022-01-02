package entity

import "github.com/aokuyama/go-stock_jp/value"

type ITrade interface {
	String() string
	Target() *value.SecurityCode
}

type IPosition interface {
	String() string
	Target() *value.SecurityCode
	Type() *value.PositionType
}
