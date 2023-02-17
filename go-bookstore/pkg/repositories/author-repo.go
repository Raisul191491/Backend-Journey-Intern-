package repositories

import (
	"github.com/deadking/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

type dba struct {
	DB *gorm.DB
}

func AuthorDbInstance(d *gorm.DB) IAuthorCRUD {
	db = d
	return &dba{
		DB: d,
	}
}

func (repo *dba) Create(a models.Author) (*models.Author, string) {
	db.Table("authors").Create(&a)
	return &a, "Author created, Successfully"
}

// func (repo *dba) Update(ID int, updateAuthor models.Author) (models.Author, string)

func (repo *dba) Delete(ID int) string {
	var deletedAuthor models.Author
	var deletedBooks []models.Book

	db.Where("id=?", ID).Find(&deletedAuthor)
	if deletedAuthor.AuthorName == "" || deletedAuthor.Age == 0 {
		return "Author not found to begin with"
	}
	db.Where("author_id=?", ID).Delete(&deletedBooks)
	db.Where("id=?", ID).Delete(&deletedAuthor)
	return "Successfully deleted...."
}

func (repo *dba) Get(authorID int) []models.Author {
	var authors []models.Author
	if authorID > 0 {
		db.Where("id=?", authorID).Find(&authors)
		return authors
	}
	db.Find(&authors)
	return authors
}
