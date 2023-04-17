package handlers

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/database"
	"github.com/minacio00/adopet/models"
)

func CreatePet(c *fiber.Ctx) error {
	c.Accepts("application/json")
	pet := &models.Pet{}
	err := c.BodyParser(&pet)
	if err != nil {
		println(err.Error())
		return c.SendStatus(400)
	}
	if pet.AbrigoID == 0 || pet.Nome == "" || pet.Idade == "" || pet.Imagem == "" {
		return c.Status(400).SendString("campos não podem estar vazios")
	}
	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regex.MatchString(pet.Nome) {
		return c.Status(400).SendString("Campo nome com caracteres inválidos")
	}
	// todo: checar se pet.AbrigoID se refere a um abrigo existente
	abrigo := models.Abrigo{}
	err = database.Db.Model(&models.Abrigo{}).Find(&abrigo, "id = ?", pet.AbrigoID).Error
	if err != nil {
		println(err.Error())
	}
	if abrigo.ID == 0 {
		return c.Status(400).SendString("abrigo não encontrado")
	}
	return c.Status(200).JSON(&abrigo)
}
