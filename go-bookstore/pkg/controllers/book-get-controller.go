package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var (
	defaultBookID   string = "0"
	defaultAuthorID string = "0"
)

func GetBookAnyway(w http.ResponseWriter, r *http.Request) {

	// Getting query data
	tempBookID := r.URL.Query().Get("bookId")
	tempAuthorID := r.URL.Query().Get("authorId")

	bookId := tempBookID
	if tempBookID == "" {
		bookId = defaultBookID
	}
	authorId := tempAuthorID
	if tempAuthorID == "" {
		authorId = defaultAuthorID
	}

	tempBook, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Enter a valid Book ID :", err)
	}

	tempAuthor, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Enter a valid Author ID :", err)
	}

	books := BookInt.Get(int(tempBook), int(tempAuthor))

	res, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
