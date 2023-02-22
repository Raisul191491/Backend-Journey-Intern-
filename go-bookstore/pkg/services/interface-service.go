package services

import (
	"github.com/deadking/go-bookstore/pkg/models"
)

var BookInterface models.IBookCRUD
var AuthorInterface models.IAuthorCRUD

func BookInterfaceInstance(bookInt models.IBookCRUD) {
	BookInterface = bookInt
}

func AuthorInterfaceInstance(authorInt models.IAuthorCRUD) {
	AuthorInterface = authorInt
}
