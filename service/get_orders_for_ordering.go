package service

import (
	"github.com/aokuyama/go-stock_jp/model/common"
	"github.com/aokuyama/go-stock_jp/model/order"
)

type GetOrdersForOrdering struct {
	Repository order.OrderRepository
}

func NewGetOrdersForOrdering(repository order.OrderRepository) *GetOrdersForOrdering {
	return &GetOrdersForOrdering{
		Repository: repository,
	}
}

func (s *GetOrdersForOrdering) GetOrders(today *common.Date) (*order.Collection, error) {
	orders, err := s.Repository.LoadByDate(today)
	if err != nil {
		return nil, err
	}
	return orders.PickupOrderingOrders(), nil
}
