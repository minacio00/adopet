package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Adocao struct {
	gorm.Model
	PetID   uint `json:"animal"`
	Pet     Pet
	TutorID uint `json:"tutor"`
	Tutor   Tutor
	Data    time.Time `json:"data"`
}

func (a *Adocao) Validate() error {
	if a.PetID == 0 {
		return fmt.Errorf("animal inválido")
	}
	if a.TutorID == 0 {
		return fmt.Errorf("tutor inválido")
	}
	// layout := "2006-01-02 15:04:05.000000-07"
	// if _, err := time.Parse(layout, a.Data.String()); err != nil {
	// 	return fmt.Errorf("data inválida")
	// }
	if a.Data.Equal(time.Time{}) {
		return fmt.Errorf("data inválida")
	}
	return nil
}
