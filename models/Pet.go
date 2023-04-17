package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Nome      string `json:"nome" gorm:"not null"`
	Descricao string `json:"descricao"`
	AbrigoID  uint   `json:"abrigo"`
	Adotado   bool   `json:"adotado"`
	Idade     string `json:"idade" gorm:"not null"`
	Imagem    string `json:"imagem"`
}
