package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/haserta98/go-rmq-learning/shared"
	"github.com/haserta98/go-rmq-learning/shared/messaging"
	"log"
)

func main() {
	app := fiber.New()

	client := messaging.NewRMQClient("amqp://guest:guest@localhost:5672/")
	var producer = NewOrderCreateProducer(client)
	app.Get("/", func(ctx fiber.Ctx) error {
		producer.CreateOrder(&shared.Order{
			Price:  100.0,
			UserID: 5,
		})
		return ctx.SendString("Hello, World!")
	})

	app.Post("/order", func(ctx fiber.Ctx) error {
		var order shared.Order
		if err := json.Unmarshal(ctx.Body(), &order); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Invalid order data")
		}
		if err := producer.CreateOrder(&order); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to create order")
		}
		// Assuming the order is created successfully, you can return the order details
		return ctx.Status(fiber.StatusCreated).JSON(order)
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
