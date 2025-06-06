package events

import "github.com/haserta98/go-rmq/shared"

type OrderCreateEvent struct {
	shared.Event
	name  string
	order *shared.Order
}

func NewOrderCreateEvent(order *shared.Order) *OrderCreateEvent {
	return &OrderCreateEvent{
		name:  "order.created",
		order: order,
	}
}

func (e *OrderCreateEvent) Name() string {
	return e.name
}

func (e *OrderCreateEvent) Data() any {
	return e.order
}
