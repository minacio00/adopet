package handlers

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/database"
	"github.com/minacio00/adopet/models"
)

func isValidImageURL(urlStr string) bool {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	ext := strings.ToLower(parsedURL.Path[len(parsedURL.Path)-4:])
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return false
	}

	return true
}

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

	//checa se o abrigo existe
	abrigo := models.Abrigo{}
	err = database.Db.Model(&models.Abrigo{}).Find(&abrigo, "id = ?", pet.AbrigoID).Error
	if err != nil {
		println(err.Error())
	}
	if abrigo.ID == 0 {
		return c.Status(400).SendString("abrigo não encontrado")
	}
	if !isValidImageURL(pet.Imagem) {
		return c.Status(400).SendString("url de imagem inválida")
	}
	err = database.Db.Save(&pet).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(200).JSON(&pet)
}

func GetAllPets(c *fiber.Ctx) error {
	c.Accepts("application/json")
	pets := &[]models.Pet{}
	err := database.Db.Find(pets).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(200).JSON(pets)
}
