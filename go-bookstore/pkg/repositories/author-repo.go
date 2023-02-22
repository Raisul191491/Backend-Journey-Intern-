package repositories

import (
	"errors"

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

func (repo *dba) Create(a models.Author) (*types.ResponseAuthor, error) {
	responseAuthor := types.ResponseAuthor{}
	if err := db.Table("authors").Create(&a).Error; err != nil {
		return &responseAuthor, err
	}
	responseAuthor = types.ResponseAuthor(a)
	return &responseAuthor, nil
}

// func (repo *dba) Update(ID int, updateAuthor models.Author) (models.Author, string)

func (repo *dba) Delete(ID int) error {

	tempAuthors := repo.Get(ID)
	if len(tempAuthors) == 0 {
		return errors.New("no Author Found")
	}
	var deletedAuthor models.Author
	var deletedBooks []models.Book

	// if deletedAuthor.AuthorName == "" || deletedAuthor.Age == 0 {
	// 	return ErrorResponse.Error("Author not found ")
	// }
	db.Where("author_id=?", ID).Delete(&deletedBooks)
	if err := db.Where("id=?", ID).Delete(&deletedAuthor).Error; err != nil {
		return err
	}
	return nil
}

func Get(ID int) {
	panic("unimplemented")
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
