package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
	"github.com/gorilla/mux"
)


func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db = config.GetDB()
	finalMsg := types.CustomBookResponse{}
	var book, deletedBook models.Book

	// Getting query data
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Deleting row
	db.Where("ID=?", ID).Find(&deletedBook)
	if deletedBook.Name == "" || deletedBook.Publication == "" {
		finalMsg.Msg = "Book not found to begin with"
	} else {
		db.Where("ID=?", ID).Delete(&book)
		finalMsg.Msg = "Successfully deleted...."
	}
	finalMsg.Content = deletedBook

	res, _ := json.Marshal(finalMsg)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
