package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	book := models.Book{}
	finalMsg := types.CustomBookResponse{}
	db = config.GetDB()

	//  Create row
	json.NewDecoder(r.Body).Decode(&book)
	err := book.Validate()
	if err == nil {
		db.Table("books").Create(&book)
		finalMsg.Msg = "Book created, Successfully"
	} else {
		finalMsg.Msg = err.Error()
	}
	finalMsg.Content = book

	res, err := json.Marshal(finalMsg)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
