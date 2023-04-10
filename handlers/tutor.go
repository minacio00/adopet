package handlers

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

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
		return c.Status(400).JSON(struct{ Body string }{
			Body: "nome, email e password nao podem ser nulos"})
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println("Error:", err)
	}
	user.Password = string(bytes)

	result := database.Db.Save(user)
	if result.Error != nil {
		return c.Status(fiber.StatusExpectationFailed).SendString(result.Error.Error())
	}

	return c.Status(200).JSON(struct{ Message string }{Message: "tutor created"})
}

func ListTutors(c *fiber.Ctx) error {
	c.Accepts("application/json")
	// users := []models.Tutor{}
	users := []struct {
		Nome     string `json:"nome"`
		Foto     string `json:"foto"`
		Telefone string `json:"telefone"`
		Cidade   string `json:"cidade"`
		Sobre    string `json:"sobre"`
		Email    string
	}{}
	database.Db.Model(&models.Tutor{}).Select("nome, foto, telefone, cidade, sobre, email").Find(&users)
	if len(users) == 0 {
		return c.Status(200).JSON(struct{ Message string }{Message: "Nenhum tutor cadastrado"})
	}
	return c.Status(200).JSON(&users)
}

func UpatadeTutor(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, _ := c.ParamsInt("id", -1)

	if id == -1 {
		return c.Status(400).SendString("informe um id")
	}
	var updatades map[string]interface{}
	var tutor models.Tutor
	c.BodyParser(&updatades) // campos recebidos na requisição

	database.Db.Find(&tutor, "id = ?", id)
	if tutor.ID == 0 {
		return c.Status(400).JSON(struct{ Message string }{Message: "Tutor não encontrado"})
	}
	database.Db.Model(&tutor).Updates(&updatades)

	return c.Status(200).JSON(&tutor)
}
func DeleteTutor(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, _ := c.ParamsInt("id", -1)
	if id == -1 {
		return c.Status(400).SendString("informe um id")
	}
	tutor := models.Tutor{}
	err := database.Db.Delete(&tutor, "id = ?", id).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(200).SendString("Tutor apagado com sucesso!")
}

func FindTutor(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, _ := c.ParamsInt("id", -1)
	if id == -1 {
		return c.Status(400).SendString("informe um id")
	}

	tutor := models.Tutor{}
	err := database.Db.Model(models.Tutor{}).Select("nome, foto, telefone, cidade, sobre, email").First(&tutor, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(&tutor)
}
