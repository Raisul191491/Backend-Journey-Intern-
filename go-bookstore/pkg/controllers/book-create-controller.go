package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	book := models.Book{}
	finalMsg := types.CustomBookResponse{}
	db = config.GetDB()

	json.NewDecoder(r.Body).Decode(&book)
	err := book.Validate()
	if err == nil {
		db.Table("books").Create(&book)
		finalMsg.Msg = "Successfully created"
	} else {
		finalMsg.Msg = err.Error()
	}
	finalMsg.Content = book

	res, _ := json.Marshal(finalMsg)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
