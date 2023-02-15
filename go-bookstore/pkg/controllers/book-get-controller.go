package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/models"
)

var (
	BasicID string = "0"
)

func GetBookAnyway(w http.ResponseWriter, r *http.Request) {
	db = config.GetDB()
	var books []models.Book

	// Getting query data
	q := r.URL.Query().Get("bookId")
	bookId := q
	if q == "" {
		bookId = BasicID
	}
	temp, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Enter a valid Book ID :", err)
	}

	// Retrieve data
	if temp > 0 {
		db.Where("ID=?", temp).Find(&books)
	} else {
		db.Find(&books)
	}

	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
