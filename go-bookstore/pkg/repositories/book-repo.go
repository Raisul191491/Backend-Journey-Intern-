package repositories

import (
	"github.com/deadking/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

var db *gorm.DB

type dbs struct {
	DB *gorm.DB
}

func BookDbInstance(d *gorm.DB) IBookCRUD {
	db = d
	return &dbs{
		DB: d,
	}
}

func (repo *dbs) Create(b models.Book) (*models.Book, string) {
	author := models.Author{}
	if err := db.
		Table("authors").
		Where("id=?", b.AuthorID).
		First(&author).
		Error; err == nil {
		db.Table("books").Create(&b)
		return &b, "Book created, Successfully"
	}
	return nil, "Author does not exist"
}

func (repo *dbs) Delete(ID int) string {
	var deletedBook models.Book

	db.Where("ID=?", ID).Find(&deletedBook)
	if deletedBook.Name == "" || deletedBook.Publication == "" {
		return "Book not found to begin with"
	}
	db.Where("ID=?", ID).Delete(&deletedBook)
	return "Successfully deleted...."
}

func (repo *dbs) Update(ID int, updateBook models.Book) (models.Book, string) {
	var book models.Book
	db.Where("ID=?", ID).Find(&book)
	if book.Name == "" || book.Publication == "" {
		return models.Book{}, "Book not found"
	}

	// Update or reject update
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	err := book.Validate()
	if err == nil {
		db.Save(&book)
		return book, "Successfully updated"
	}
	return book, err.Error()
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
