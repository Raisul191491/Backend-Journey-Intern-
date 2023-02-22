package services

import (
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
)

func CreateAuthorService(a models.Author) (*types.ResponseAuthor, error) {
	finalMsg := types.CustomAuthorResponse{}
	finalMsg.Content, finalMsg.ErrorMsg = AuthorInterface.Create(a)
	return finalMsg.Content, finalMsg.ErrorMsg
}

func DeleteAuthorService(ID int) error {
	msg := AuthorInterface.Delete(int(ID))
	return msg
}

func GetAuthorService(authorID int) []types.ResponseAuthor {
	authorlist := AuthorInterface.Get(authorID)
	return authorlist
}
