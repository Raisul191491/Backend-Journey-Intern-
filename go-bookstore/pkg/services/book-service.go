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
	var responseBooks []types.ResponseBook
	booklist := BookInterface.Get(bookId, authorId)
	for _, val := range booklist {
		responseBooks = append(responseBooks, types.ResponseBook{
			ID:          val.ID,
			Name:        val.Name,
			Publication: val.Publication,
			AuthorID:    val.AuthorID,
			Author: types.ResponseAuthor{
				AuthorName: val.Author.AuthorName,
				Age:        val.Author.Age,
			},
		})
	}
	return responseBooks
}

func UpdateBookService(updateBook models.Book) (*types.ResponseBook, error) {
	finalMsg := types.CustomBookResponse{}
	finalMsg.Content, finalMsg.ErrorMsg = BookInterface.Update(updateBook)
	return finalMsg.Content, finalMsg.ErrorMsg
}
