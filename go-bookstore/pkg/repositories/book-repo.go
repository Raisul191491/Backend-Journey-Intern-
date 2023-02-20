package repositories

import (
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

func (repo *dbs) Create(b models.Book) (*types.ResponseBook, string) {
	author := models.Author{}
	responseBook := types.ResponseBook{}
	if err := db.
		Table("authors").
		Where("id=?", b.AuthorID).
		First(&author).
		Error; err == nil {
		db.Table("books").Create(&b)
		return &responseBook, "Book created, Successfully"
	}
	return nil, "Author does not exist"
}

func (repo *dbs) Delete(ID int) string {
	var deletedBook models.Book

	db.Where("ID=?", ID).Find(&deletedBook)
	if err := db.
		Table("books").
		Where("id=?", ID).
		First(&deletedBook).
		Error; err != nil {
		return "Book not found to begin with..."
	}
	db.Where("ID=?", ID).Delete(&deletedBook)
	return "Successfully deleted...."
}

func (repo *dbs) Update(ID int, updateBook models.Book) (*types.ResponseBook, string) {
	var book models.Book
	var responseBook types.ResponseBook
	db.Where("ID=?", ID).Find(&book)
	if book.Name == "" || book.Publication == "" {
		return &types.ResponseBook{}, "Book not found"
	}

	// Update or reject update
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	if updateBook.AuthorID != 0 {
		book.AuthorID = updateBook.AuthorID
	}
	err := book.Validate()
	if err == nil {
		db.Save(&book)
		responseBook = types.ResponseBook{
			ID:          book.ID,
			Name:        book.Name,
			Publication: book.Publication,
			AuthorID:    book.AuthorID,
			Author:      types.ResponseAuthor(book.Author),
		}
		return &responseBook, "Successfully updated"
	}
	return &responseBook, err.Error()
}

func (repo *dbs) Get(bookID, authorID int) []types.ResponseBook {
	var books []models.Book
	var responseBooks []types.ResponseBook

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
	for _, val := range books {
		responseBooks = append(responseBooks, types.ResponseBook(types.ResponseBook{
			ID:          val.ID,
			Name:        val.Name,
			Publication: val.Publication,
			AuthorID:    val.AuthorID,
			Author:      types.ResponseAuthor(val.Author),
		}))
	}
	return responseBooks
}
