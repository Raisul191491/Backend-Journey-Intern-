package repositories

import (
	"errors"

	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
	"gorm.io/gorm"
)

var db *gorm.DB

type dbs struct {
	DB *gorm.DB
}

func BookDbInstance(d *gorm.DB) models.IBookCRUD {
	db = d
	return &dbs{
		DB: d,
	}
}

func (repo *dbs) Create(b models.Book) (*types.ResponseBook, error) {
	author := models.Author{}
	err := db.
		Table("authors").
		Where("id=?", b.AuthorID).
		First(&author).
		Error
	if err == nil {
		db.Table("books").Create(&b)
		responseBook := types.ResponseBook{
			ID:          b.ID,
			Name:        b.Name,
			Publication: b.Publication,
			AuthorID:    b.AuthorID,
			Author:      types.ResponseAuthor(b.Author),
		}
		return &responseBook, nil
	}
	return nil, err
}

func (repo *dbs) Delete(ID int) error {
	var deletedBook models.Book

	tempBooks := repo.Get(ID, 0)
	if len(tempBooks) == 0 {
		return errors.New("deletion error")
	}
	db.Where("ID=?", ID).Delete(&deletedBook)
	return nil
}

func (repo *dbs) Get(bookID, authorID int) []models.Book {
	var books []models.Book

	if bookID > 0 && authorID > 0 {
		db.
			Joins("Author").
			Where("`books`.`id`=? AND `author_id`=?", bookID, authorID).
			Find(&books)
	} else if bookID > 0 {
		db.Joins("Author").Where("`books`.`id`=?", bookID).Find(&books)
	} else if authorID > 0 {
		db.Joins("Author").Where("author_id=?", authorID).Find(&books)
	} else if bookID == 0 && authorID == 0 {
		db.Joins("Author").Find(&books)
	}
	return books
}

func (repo *dbs) Update(book models.Book) (*types.ResponseBook, error) {
	if err := db.Save(&book).Error; err != nil {
		return nil, err
	}
	responseBook := types.ResponseBook{
		ID:          book.ID,
		Name:        book.Name,
		Publication: book.Publication,
		AuthorID:    book.AuthorID,
		Author: types.ResponseAuthor{
			AuthorName: book.Author.AuthorName,
			Age:        book.Author.Age,
		},
	}
	return &responseBook, nil
}
