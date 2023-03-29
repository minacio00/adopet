package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/database"
	"github.com/minacio00/adopet/handlers"
)

func main() {
	app := fiber.New()
	database.Connectdb()

	app.Post("/tutores", handlers.CreateTutor)
	app.Listen(":8080")
}
