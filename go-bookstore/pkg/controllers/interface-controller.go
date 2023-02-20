package controllers

import (
	"github.com/deadking/go-bookstore/pkg/models"
)

var BookInt models.IBookCRUD
var AuthorInt models.IAuthorCRUD

func BookInterfaceInstance(bookInt models.IBookCRUD) {
	BookInt = bookInt
}

func AuthorInterfaceInstance(authorInt models.IAuthorCRUD) {
	AuthorInt = authorInt
}
