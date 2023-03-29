package models

import "gorm.io/gorm"

type Tutor struct {
	gorm.Model
	Nome     string `json:"nome" gorm:"not null"`
	Foto     string `json:"foto"`
	Telefone string `json:"telefone"`
	Cidade   string `json:"cidade"`
	Sobre    string `json:"sobre"`
	Email    string `gorm:"not null; unique"`
	Password string `gorm:"not null"`
}
