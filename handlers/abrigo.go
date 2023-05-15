package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/database"
	"github.com/minacio00/adopet/models"
	"gorm.io/gorm"
)

func CreateAbrigo(c *fiber.Ctx) error {
	c.Accepts("application/json")
	abrigo := &models.Abrigo{}
	err := json.Unmarshal(c.Body(), abrigo)
	if err != nil {
		println(err.Error())
	}
	err = abrigo.Validate()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = database.Db.Create(&abrigo).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(200).JSON(&abrigo)
}

func GetAbrigo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		println(err.Error())
	}
	abrigo := &models.Abrigo{}
	err = database.Db.Preload("Pets").First(&abrigo, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("não encontrado")
	}
	return c.JSON(&abrigo)
}

func GetAbrigos(c *fiber.Ctx) error {
	abrigos := &[]models.Abrigo{}
	err := database.Db.Preload("Pets").Find(&abrigos).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(fiber.StatusAccepted).JSON(&abrigos)
}
func DeleteAbrigo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("id inválido")
	}
	abrigo := &models.Abrigo{}
	err = database.Db.First(&abrigo, id).Error
	if err != nil {
		println(err.Error())
	}

	err = database.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Select("Pets").Delete(&abrigo).Error
	})
	if err != nil {
		println(err.Error())
	}
	return c.JSON(&abrigo)

}

func UpdateAbrigo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("id inválido")
	}

	body := &models.Abrigo{}
	err = c.BodyParser(&body)
	if err != nil {
		println(err.Error())
	}
	err = body.Validate()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	abrigo := &models.Abrigo{}
	err = database.Db.First(&abrigo, "id = ?", id).Error
	if err != nil {
		println(err.Error())
	}
	if abrigo.ID == 0 {
		return c.Status(400).SendString("Abrigo não encontrado")
	}
	err = json.Unmarshal(c.Body(), &abrigo)
	if err != nil {
		println(err.Error())
	}
	err = database.Db.Model(&abrigo).Updates(&abrigo).Error
	if err != nil {
		println(err.Error())
	}
	return c.Status(200).JSON(&abrigo)
}
