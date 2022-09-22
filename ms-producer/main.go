package main

import (
	"github.com/anarkigotic/fiber/db"
	"github.com/anarkigotic/fiber/send"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func main() {

	app := fiber.New()

	producer := app.Group("/send")
	producer.Post("", func(c *fiber.Ctx) error {
		p := new(User)

		if err := c.BodyParser(p); err != nil {
			return err
		}
		go db.WriteFile("postgress.txt", 5, string(p.Name))
		go send.SendMessageALl(p.Name)

		return c.SendString("ok")
	})
	app.Listen(":8080")

}
