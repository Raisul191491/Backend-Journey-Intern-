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
	DefaultBookID   string = "0"
	DefaultAuthorID string = "0"
)

func GetBookAnyway(w http.ResponseWriter, r *http.Request) {
	db = config.GetDB()
	var books []models.Book

	// Getting query data
	tempBookID := r.URL.Query().Get("bookId")
	tempAuthorID := r.URL.Query().Get("authorId")

	bookId := tempBookID
	if tempBookID == "" {
		bookId = DefaultBookID
	}
	authorId := tempAuthorID
	if tempAuthorID == "" {
		authorId = DefaultAuthorID
	}

	tempBook, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Enter a valid Book ID :", err)
	}

	tempAuthor, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Enter a valid Book ID :", err)
	}

	// Retrieve data by book id
	if tempBook > 0 {
		db.Where("ID=?", tempBook).Find(&books)
	} else {
		db.Find(&books)
	}

	// Retrieve data by author id
	if tempAuthor > 0 {
		db.Where("author_id=?", tempAuthor).Find(&books)
	} else {
		db.Find(&books)
	}

	res, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
