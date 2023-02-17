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

func CreateBook(w http.ResponseWriter, r *http.Request) {

	book := models.Book{}
	finalMsg := types.CustomBookResponse{}

	json.NewDecoder(r.Body).Decode(&book)
	err := book.Validate()
	if err == nil {
		finalMsg.Content, finalMsg.Msg = BookInt.Create(book)
	} else {
		finalMsg.Content, finalMsg.Msg = &book, err.Error()
	}

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
	w.Header().Set("Content-Type", "pkglication/json")
	// Getting query data
	bookId := r.URL.Query().Get("bookId")
	authorId := r.URL.Query().Get("authorId")

	tempBook, errBook := strconv.ParseInt(bookId, 0, 0)

	tempAuthor, errAuthor := strconv.ParseInt(authorId, 0, 0)

	books := BookInt.Get(int(tempBook), int(tempAuthor))
	res, err := json.Marshal(books)

	if len(books) == 0 {
		res, err = json.Marshal("Msg : No books found")
	} else if bookId != "" && errBook != nil {
		res, err = json.Marshal("Msg : Invalid book Id")
		w.WriteHeader(http.StatusBadRequest)
	} else if authorId != "" && errAuthor != nil {
		res, err = json.Marshal("Msg : Invalid author Id")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
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
