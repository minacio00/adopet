package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/database"
)

func main() {
	app := fiber.New()
	database.Connectdb()

	app.Post("/tutores", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.SendStatus(200)
	})

}
