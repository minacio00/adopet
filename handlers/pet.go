package handlers

import (
	"encoding/json"
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
	pets := &[]models.Pet{}
	err := database.Db.Find(&pets).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(200).JSON(pets)
}
func GetPet(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		println(err.Error())
	}
	pet := &models.Pet{}
	err = database.Db.Find(&pet, "id = ?", id).Error
	if err != nil {
		println(err.Error())
	}
	if pet.ID == 0 {
		return c.Status(404).SendString("pet não encontrado")
	}
	return c.Status(200).JSON(&pet)
}

func UpdatePet(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body, pet := &models.Pet{}, &models.Pet{}
	err := c.BodyParser(&body)
	if err != nil {
		println(err.Error())
	}
	// checar se existe no banco um pet com o id dado
	if body.ID == 0 {
		return c.Status(400).SendString("id inválido")
	}
	err = database.Db.Find(&pet, "id = ?", body.ID).Error
	if err != nil {
		println(err.Error())
	}
	if pet.ID == 0 {
		return c.Status(404).SendString("pet não encontrado")
	}

	if body.AbrigoID == 0 || body.Nome == "" || body.Idade == "" || body.Imagem == "" {
		return c.Status(400).SendString("campos não podem estar vazios")
	}
	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regex.MatchString(body.Nome) {
		return c.Status(400).SendString("Campo nome com caracteres inválidos")
	}

	//checa se o abrigo existe
	abrigo := models.Abrigo{}
	err = database.Db.Model(&models.Abrigo{}).Find(&abrigo, "id = ?", body.AbrigoID).Error
	if err != nil {
		println(err.Error())
	}
	if abrigo.ID == 0 {
		return c.Status(400).SendString("abrigo não encontrado")
	}
	if !isValidImageURL(pet.Imagem) {
		return c.Status(400).SendString("url de imagem inválida")
	}
	json.Unmarshal(c.Body(), &pet)
	err = database.Db.Save(&pet).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(200).JSON(&pet)

}

func DeletePet(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := c.ParamsInt("id")
	if err != nil {
		println(err.Error())
	}
	if id == 0 {
		return c.Status(404).SendString("pet não encontrado")
	}
	pet := models.Pet{}
	err = database.Db.Find(&pet, "id = ?", id).Error
	if err != nil {
		println(err.Error())
	}
	if pet.ID == 0 {
		return c.Status(404).SendString("pet não encontrado")
	}

	err = database.Db.Delete(&pet, "id = ?", id).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(200).SendString("pet deletado com sucesso")
}
