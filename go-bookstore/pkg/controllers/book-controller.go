package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/services"
	"github.com/deadking/go-bookstore/pkg/types"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")

	book := models.Book{}
	finalMsg := types.CustomBookResponse{}

	json.NewDecoder(r.Body).Decode(&book)
	err := book.Validate()
	if err == nil {
		finalMsg.Content, finalMsg.Msg = services.CreateBookService(book)
	} else {
		finalMsg.Content, finalMsg.Msg = nil, err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	res, err := json.Marshal(finalMsg)
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")

	finalMsg := types.CustomDeleteResponse{}
	// Getting query data
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		finalMsg.Msg = "Enter valid book ID"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		finalMsg.Msg = services.DeleteBookService(int(ID))
		w.WriteHeader(http.StatusOK)
	}
	res, err := json.Marshal(finalMsg)
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.Write(res)
}

func GetBookAnyway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	// Getting query data
	bookId := r.URL.Query().Get("bookId")
	authorId := r.URL.Query().Get("authorId")

	tempBook, errBook := strconv.ParseInt(bookId, 0, 0)

	tempAuthor, errAuthor := strconv.ParseInt(authorId, 0, 0)

	books := services.GetBookService(int(tempBook), int(tempAuthor))
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
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	updateBook := models.Book{}
	finalMsg := types.CustomBookResponse{}

	json.NewDecoder(r.Body).Decode(&updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		finalMsg.Msg = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		finalMsg.Content, finalMsg.Msg = services.UpdateBookService(int(ID), updateBook)
		w.WriteHeader(http.StatusAccepted)
	}
	res, err := json.Marshal(finalMsg)
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.Write(res)
}
