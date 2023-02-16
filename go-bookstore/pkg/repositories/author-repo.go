package repositories

import "github.com/deadking/go-bookstore/pkg/models"

func CreateAuthor(a models.Author) (models.Author, string) {
	err := a.Validate()
	if err == nil {
		db.Table("books").Create(&a)
		return a, "Author created, Successfully"
	}
	return a, err.Error()
}