package services

import (
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
)

func CreateBookService(b models.Book) (*types.ResponseBook, error) {
	finalMsg := types.CustomBookResponse{}
	finalMsg.Content, finalMsg.ErrorMsg = BookInterface.Create(b)
	return finalMsg.Content, finalMsg.ErrorMsg
}

func DeleteBookService(ID int) error {
	msg := BookInterface.Delete(ID)
	return msg
}

func GetBookService(bookId, authorId int) []types.ResponseBook {
	booklist := BookInterface.Get(bookId, authorId)
	return booklist
}

func UpdateBookService(updateBook models.Book) (*types.ResponseBook, error) {
	finalMsg := types.CustomBookResponse{}
	finalMsg.Content, finalMsg.ErrorMsg = BookInterface.Update(updateBook)
	return finalMsg.Content, finalMsg.ErrorMsg
}
