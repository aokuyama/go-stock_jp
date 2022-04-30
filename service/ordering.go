package service

import (
	"github.com/aokuyama/go-generic_subdomains/event"
	"github.com/aokuyama/go-stock_jp/model/order"
)

type Ordering struct {
	Repository order.OrderRepository
}

func NewOrdering(repo order.OrderRepository) *Ordering {
	return &Ordering{
		Repository: repo,
	}
}

func (s *Ordering) OrderingOrders(orders *order.Collection, p *event.Publisher) (*order.Collection, error) {
	var odrngs []*order.Order
	for _, o := range *orders {
		odrng, err := s.Ordering(o, p)
		if err != nil {
			return nil, err
		}
		odrngs = append(odrngs, odrng)
	}
	return order.NewCollection(odrngs...)
}

func (s *Ordering) Ordering(o *order.Order, p *event.Publisher) (*order.Order, error) {
	err := p.Publish(order.NewOrderingEvent(o))
	if err != nil {
		return nil, err
	}
	odrng := o.Ordering()
	err = s.Repository.Save(odrng)
	if err != nil {
		return nil, err
	}
	return odrng, nil
}
