package models

import "gorm.io/gorm"

type Abrigo struct {
	gorm.Model
	Pets []Pet
}
