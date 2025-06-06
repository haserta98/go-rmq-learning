package main

import (
	"github.com/haserta98/go-rmq/order/events"
	"github.com/haserta98/go-rmq/shared"
	"github.com/haserta98/go-rmq/shared/messaging"
	ampq "github.com/rabbitmq/amqp091-go"
)

type OrderCreateProducer struct {
	conn  *messaging.RMQClient
	queue ampq.Queue
}

func NewOrderCreateProducer(conn *messaging.RMQClient) *OrderCreateProducer {
	ch, err := conn.Conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		"order.created",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return &OrderCreateProducer{conn: conn, queue: queue}
}

func (producer *OrderCreateProducer) CreateOrder(order *shared.Order) error {
	var event shared.Event = events.NewOrderCreateEvent(order)
	producer.conn.Publish(event)
	return nil
}
