package models

import "gorm.io/gorm"

type Abrigo struct {
	gorm.Model
	Pets   []Pet  `json:"pets" gorm:"constraint:OnDelete:CASCADE"`
	Cidade string `json:"cidade" gorm:"not null"`
	Uf     string `json:"uf" gorm:"not null"`
	Nome   string `json:"nome" gorm:"not null"`
}
