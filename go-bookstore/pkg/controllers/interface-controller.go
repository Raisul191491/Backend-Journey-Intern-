package controllers

import "github.com/deadking/go-bookstore/pkg/repositories"

var BookInt repositories.IBookCRUD
var AuthorInt repositories.IAuthorCRUD

func BookInterfaceInstance(bookInt repositories.IBookCRUD) {
	BookInt = bookInt
}

func AuthorInterfaceInstance(authorInt repositories.IAuthorCRUD) {
	AuthorInt = authorInt
}
