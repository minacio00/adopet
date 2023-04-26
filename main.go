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
	app.Get("/tutores", handlers.ListTutors)
	app.Patch("/tutores/:id", handlers.UpatadeTutor)
	app.Delete("/tutores/:id", handlers.DeleteTutor)
	app.Get("/tutores/:id", handlers.FindTutor)

	app.Post("/pet", handlers.CreatePet)
	app.Get("/pets", handlers.GetAllPets)
	app.Get("/pet/:id", handlers.GetPet)
	app.Put("/pet", handlers.UpdatePet)
	app.Delete("/pet/:id", handlers.DeletePet)
	app.Listen(":8080")
}
