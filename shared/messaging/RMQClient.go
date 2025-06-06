package messaging

import (
	"encoding/json"
	"github.com/haserta98/go-rmq/shared"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RMQClient struct {
	Conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRMQClient(url string) *RMQClient {
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return &RMQClient{
		Conn:    conn,
		channel: ch,
	}
}

func (client *RMQClient) Publish(event shared.Event) {
	name := event.Name()
	data := event.Data()
	if data == nil {
		panic("data cannot be nil")
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	_err := client.channel.Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        byteData,
		},
	)
	if _err != nil {
		return
	}
}
