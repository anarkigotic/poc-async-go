package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	EventsAS()
	app := fiber.New()

	producer := app.Group("/consumer")
	producer.Get("", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
	app.Listen(":8081")

}
