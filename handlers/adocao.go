package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/database"
	"github.com/minacio00/adopet/models"
	"gorm.io/gorm"
)

type body struct {
	PetID   uint `json:"animal"`
	TutorID uint `json:"tutor"`
}

func CreateAdocao(c *fiber.Ctx) error {
	c.Accepts("application/json")

	//informa o formato de timestamp que será recebido por time.Parse
	layout := "2006-01-02 15:04:05.000000-07"
	dateString := struct {
		Data string `json:"data"`
	}{}
	err := c.BodyParser(&dateString) // recebendo como string pois time.Time espera um formato diferente do que está sendo recebido
	if err != nil {
		return err
	}
	parsedDate, err := time.Parse(layout, dateString.Data)
	if err != nil {
		return err
	}

	data := body{}
	err = c.BodyParser(&data)
	if err != nil {
		return err
	}

	adocao := &models.Adocao{}
	adocao.PetID = data.PetID
	adocao.TutorID = data.TutorID
	adocao.Data = parsedDate
	err = adocao.Validate()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = database.Db.Transaction(func(tx *gorm.DB) error {
		pet := &models.Pet{}
		tutor := &models.Tutor{}
		err := tx.First(&pet, "id = ?", &adocao.PetID).Error
		if err != nil {
			return err
		}

		err = tx.First(&tutor, "id = ?", &adocao.TutorID).Error
		if err != nil {
			return err
		}

		pet.Adotado = true
		tx.Save(&pet)
		tx.Create(&adocao)
		return nil
	})
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).JSON(adocao)
}
