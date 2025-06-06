package consumers

import (
	"encoding/json"
	"github.com/haserta98/go-rmq-learning/shared"
	"github.com/haserta98/go-rmq-learning/shared/messaging"
	"log"
)

type OrderCreateConsumer struct {
	client   *messaging.RMQClient
	callback func(message *shared.Order) error
}

func NewOrderCreateConsumer(client *messaging.RMQClient, callback func(message *shared.Order) error) *OrderCreateConsumer {
	return &OrderCreateConsumer{
		client:   client,
		callback: callback,
	}
}

func (consumer *OrderCreateConsumer) Consume() {
	connection := consumer.client.Conn
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"order.created",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			var data *shared.Order
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Printf("Error parsing message: %s", err)
				d.Nack(false, false)
				continue
			}
			err = consumer.callback(data)
			if err != nil {
				log.Printf("Error processing message: %s", err)
				d.Nack(false, false)
				continue
			}
			err = d.Ack(false)
			if err != nil {
				log.Printf("Error acknowledging message: %s", err)
				continue
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
