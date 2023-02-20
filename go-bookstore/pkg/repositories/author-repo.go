package repositories

import (
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
	"gorm.io/gorm"
)

type dba struct {
	DB *gorm.DB
}

func AuthorDbInstance(d *gorm.DB) models.IAuthorCRUD {
	db = d
	return &dba{
		DB: d,
	}
}

func (repo *dba) Create(a models.Author) (*types.ResponseAuthor, string) {
	responseAuthor := types.ResponseAuthor{}
	db.Table("authors").Create(&a)
	responseAuthor = types.ResponseAuthor(a)
	return &responseAuthor, "Author created, Successfully"
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

func (repo *dba) Get(authorID int) []types.ResponseAuthor {
	var authors []models.Author
	var responseAuthors []types.ResponseAuthor
	if authorID > 0 {
		db.Where("id=?", authorID).Find(&authors)
	} else {
		db.Find(&authors)
	}
	for _, val := range authors {
		responseAuthors = append(responseAuthors, types.ResponseAuthor(val))
	}
	return responseAuthors
}
