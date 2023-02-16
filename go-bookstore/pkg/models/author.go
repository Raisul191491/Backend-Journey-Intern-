package models

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Author struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true"`
	AuthorName string `json:"author_name"`
	Age        int    `json:"age"`
}

func ageValidate(a int) validation.RuleFunc {
	return func(value interface{}) error {
		age, _ := value.(int)
		if age < 12 || age > 130 {
			return errors.New(" must be within 12 and 130")
		}
		return nil
	}
}

func (a Author) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AuthorName, validation.Required, validation.Length(6, 150)),
		validation.Field(&a.Age, validation.By(ageValidate(a.Age))),
	)
}
