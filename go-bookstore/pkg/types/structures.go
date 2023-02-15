package types

import "github.com/deadking/go-bookstore/pkg/models"

type CustomBookResponse struct {
	Content models.Book
	Msg     string
}