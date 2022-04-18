package order

import "errors"

type Collection []*Order

func NewCollection(orders ...*Order) (*Collection, error) {
	c := Collection{}
	for _, o := range orders {
		if o == nil {
			return nil, errors.New("nil order")
		}
		c = append(c, o)
	}
	return &c, nil
}

func (c *Collection) PickupOrderingOrders() *Collection {
	nc := Collection{}
	for _, o := range *c {
		if o.CanBeOrdered() {
			nc = append(nc, o)
		}
	}
	return &nc
}
