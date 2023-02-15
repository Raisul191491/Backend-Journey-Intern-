package models

import validation "github.com/go-ozzo/ozzo-validation"

type Author struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true"`
	AuthorName string `json:"author_name"`
	Age        int    `json:"age"`
}

type AgeError struct{}

func (m *AgeError) Error() string {
	return "Age must be within 12 to 130"
}

func (a Author) Validate() error {
	if a.Age < 12 || a.Age > 130 {
		return &AgeError{}
	}
	return validation.ValidateStruct(&a,
		validation.Field(&a.AuthorName, validation.Required, validation.Length(6, 150)),
	)
}
