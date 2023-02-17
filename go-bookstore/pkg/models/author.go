package models

import (
	"errors"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Author struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Age        int    `json:"age,omitempty"`
}

func ageValidate(a int) validation.RuleFunc {
	return func(value interface{}) error {
		if value.(int) == 0 {
			return errors.New("enter valid age (Numerical)")
		}
		age := value.(int)
		if age < 12 || age > 130 {
			return errors.New("enter valid age(Numerical), within 12 and 130")
		}
		return nil
	}
}

func authorNameValidate(a string) validation.RuleFunc {
	return func(value interface{}) error {
		name := value.(string)
		if _, err := strconv.Atoi(name); err == nil || len(name) < 6 || len(name) > 150 {
			return errors.New("enter valid name(English) of 6 to 150 characters")
		}
		return nil
	}
}

func (a Author) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AuthorName,
			validation.By(authorNameValidate(a.AuthorName))),
		validation.Field(&a.Age,
			validation.By(ageValidate(a.Age))),
	)
}
