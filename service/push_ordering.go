package service

import (
	"github.com/aokuyama/go-stock_jp/model/order"
)

type PushOrdering struct {
	Orders     *order.Collection
	Repository order.OrderRepository
}

func NewPushOrdering(orders *order.Collection, repository order.OrderRepository) *PushOrdering {
	return &PushOrdering{
		Orders:     orders,
		Repository: repository,
	}
}
