package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
	Publication string `json:"publication"`
	AuthorID    uint   `json:"author_id"`
	Author      Author `json:"author" gorm:"foreignKey:AuthorID ;references:ID"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(1, 150)),
		validation.Field(&b.Publication, validation.Required, validation.Length(6, 50)),
	)
}
