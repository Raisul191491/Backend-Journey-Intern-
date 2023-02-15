package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := models.Book{}
	finalMsg := types.CustomBookResponse{}
	db = config.GetDB()

	// Getting query data
	json.NewDecoder(r.Body).Decode(&updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		finalMsg.Msg = err.Error()
	}

	// Check and retrieve the book
	var res []byte
	var book models.Book
	db.Where("ID=?", ID).Find(&book)

	// Update or reject update
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	err = book.Validate()
	if err == nil {
		db.Save(&book)
		finalMsg.Msg = "Successfully updated"
	} else {
		finalMsg.Msg = err.Error()
	}

	finalMsg.Content = book
	res, _ = json.Marshal(finalMsg)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
