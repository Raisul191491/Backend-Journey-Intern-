package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
	"github.com/gorilla/mux"
)

var (
	defaultBookID   string = "0"
	defaultAuthorID string = "0"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	book := models.Book{}
	finalMsg := types.CustomBookResponse{}

	json.NewDecoder(r.Body).Decode(&book)
	finalMsg.Content, finalMsg.Msg = BookInt.Create(book)
	res, err := json.Marshal(finalMsg)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	finalMsg := types.CustomDeleteResponse{}

	// Getting query data
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		finalMsg.Msg = err.Error()
	} else {
		finalMsg.Msg = BookInt.Delete(int(ID))
	}
	res, err := json.Marshal(finalMsg)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

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
	if len(books) == 0 {
		res, err = json.Marshal("Msg : No books found")
	}
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := models.Book{}
	finalMsg := types.CustomBookResponse{}
	// db = config.GetDB()

	// Getting query data
	json.NewDecoder(r.Body).Decode(&updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		finalMsg.Msg = err.Error()
	}

	*finalMsg.Content, finalMsg.Msg = BookInt.Update(int(ID), updateBook)
	res, err := json.Marshal(finalMsg)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
