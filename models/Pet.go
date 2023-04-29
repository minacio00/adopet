package models

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Nome      string `json:"nome" gorm:"not null"`
	Descricao string `json:"descricao"`
	AbrigoID  uint   `json:"abrigo"` //fk
	// Abrigo    Abrigo `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Adotado bool   `json:"adotado"`
	Idade   string `json:"idade" gorm:"not null"`
	Imagem  string `json:"imagem"`
}

func (p *Pet) Validate() error {
	if p.AbrigoID == 0 {
		return fmt.Errorf("abrigo nao pode ser vazio")
	}
	if p.Nome == "" {
		return fmt.Errorf("nome não pode ser vazio")
	}
	if p.Idade == "" {
		return fmt.Errorf("idade não pode ser vazio")
	}
	if p.Imagem == "" {
		return fmt.Errorf("url da imagem não pode ser vazio")
	}
	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regex.MatchString(p.Nome) {
		return fmt.Errorf("nome so pode conter letras")
	}
	return nil
}
