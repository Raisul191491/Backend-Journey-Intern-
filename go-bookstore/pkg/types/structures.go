package types

import "github.com/deadking/go-bookstore/pkg/models"

type CustomBookResponse struct {
	Content *models.Book `json:"content,omitempty"`
	Msg     string
}
type CustomAuthorResponse struct {
	Content *models.Author `json:"content,omitempty"`
	Msg     string
}

type CustomDeleteResponse struct {
	Msg string
}
