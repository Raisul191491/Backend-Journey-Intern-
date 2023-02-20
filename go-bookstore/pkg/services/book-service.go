package services

import (
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
)

func CreateBookService(b models.Book) (*types.ResponseBook, string) {
	return &types.ResponseBook{}, "Author does not exist"
}

func DeleteBookService(ID int) string {
	return "Successfully deleted...."
}

func GetBookService(bookId, authorId int) []types.ResponseBook {
	return []types.ResponseBook{}
}

func UpdateBookService(ID int, updateBook models.Book) (*types.ResponseBook, string) {
	return &types.ResponseBook{}, "Hi"
}
