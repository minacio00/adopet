package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Abrigo struct {
	gorm.Model
	Pets   []Pet  `json:"pets" gorm:"constraint:OnDelete:CASCADE"`
	Cidade string `json:"cidade" gorm:"not null"`
	Uf     string `json:"uf" gorm:"not null"`
	Nome   string `json:"nome" gorm:"not null"`
}

func (p *Abrigo) Validate() error {
	if p.Cidade == "" {
		return fmt.Errorf("cidade não pode ser vazio")
	}
	if p.Uf == "" {
		return fmt.Errorf("uf não pode ser vazio")
	}
	if p.Nome == "" {
		return fmt.Errorf("nome não pode ser vazio")
	}
	return nil
}
