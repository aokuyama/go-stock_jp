package repository

import "github.com/aokuyama/go-stock_jp/aggregate"

type Positions interface {
	BulkLoad() (*aggregate.Positions, error)
}
