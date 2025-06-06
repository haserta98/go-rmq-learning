package main

import (
	"fmt"
	"github.com/haserta98/go-rmq-learning/order/consumers"
	"github.com/haserta98/go-rmq-learning/shared"
	"github.com/haserta98/go-rmq-learning/shared/messaging"
	"github.com/haserta98/go-rmq-learning/shared/repository"
	"log"
)

var db *repository.Connection

func main() {
	fmt.Println("hello world from order")
	db = repository.NewConnection()
	db.DB.AutoMigrate(shared.Order{})

	client := initRMQ()
	initConsumer(client)

	select {}
}

func initConsumer(client *messaging.RMQClient) {
	orderCreateConsumer := consumers.NewOrderCreateConsumer(client, onOrderCreate)
	go orderCreateConsumer.Consume()
}

func initRMQ() *messaging.RMQClient {
	return messaging.NewRMQClient("amqp://guest:guest@localhost:5672/")
}

func onOrderCreate(message *shared.Order) error {
	// Here you can handle the order creation logic, such as saving it to the database
	log.Printf("Order created: %s", message.ID)
	tx := db.DB.Create(&message)
	if tx.Error != nil {
		log.Printf("Error saving order to database: %s", tx.Error)
		return tx.Error
	}
	return nil
}
