package order

import "encoding/json"

type OrderingEvent struct {
	order *Order
}

func NewOrderingEvent(o *Order) *OrderingEvent {
	return &OrderingEvent{order: o}
}

func (e *OrderingEvent) Data() string {
	s, _ := json.Marshal(e.order)
	return string(s)
}
