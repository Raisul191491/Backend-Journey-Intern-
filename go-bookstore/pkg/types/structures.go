package types

import "github.com/deadking/go-bookstore/pkg/models"

type CustomBookResponse struct {
	Content models.Book
	Msg     string
}
type CustomAuthorResponse struct {
	Content models.Author
	Msg     string
}

type CustomAuthorDeleteResponse struct {
	Content models.Author
	Books   []models.Book
	Msg     string
}
