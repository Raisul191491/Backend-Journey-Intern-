package repositories

import "github.com/deadking/go-bookstore/pkg/models"

type IBookCRUD interface {
	Create(book models.Book) (*models.Book, string)
	Update(ID int, updateBook models.Book) (models.Book, string)
	Delete(ID int) (string)
	Get(bookID, authorID int) []models.Book
}

type IAuthorCRUD interface {
	Create(book models.Author) (*models.Author, string)
	// Update(ID int, updateAuthor models.Author) (models.Author, string)
	Delete(ID int) (string)
	Get(authorID int) *[]models.Author
}
