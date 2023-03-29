package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/database"
	"github.com/minacio00/adopet/models"
)

func CreateTutor(c *fiber.Ctx) error {
	c.Accepts("application/json")
	user := &models.Tutor{}
	err := c.BodyParser(&user)
	if err != nil {
		log.Fatal(err)
		return c.SendStatus(400)
	}

	if user.Nome == "" || user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(struct{ Body string }{Body: "xiiiiii"})
	}
	hasher := sha256.New()
	hasher.Write([]byte(user.Password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	user.Password = hashedPassword
	database.Db.Save(user)

	return c.Status(200).JSON(struct{ Message string }{Message: "tutor created"})
}
