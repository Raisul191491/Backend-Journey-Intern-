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
	newBook := models.Book{
		Name:        book.Name,
		Publication: book.Publication,
		AuthorID:    book.AuthorID,
		Author: models.Author{
			ID:         book.Author.ID,
			AuthorName: book.Author.AuthorName,
			Age:        book.Author.Age,
		},
	}
	err := newBook.Validate()
	if err == nil {
		finalMsg.Content, finalMsg.ErrorMsg = services.CreateBookService(newBook)
		if finalMsg.ErrorMsg != nil {
			finalMsg.Msg = "Create request failed, Author does not exist"
			res, _ := json.Marshal(finalMsg)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}
		finalMsg.Msg = "Create request successful"
	} else {
		finalMsg.Content, finalMsg.ErrorMsg, finalMsg.Msg =
			nil, err, "Create request failed"
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

	finalMsg := types.CustomOnlyResponse{}
	// Getting query data
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil || ID < 1 {
		finalMsg.Msg = "Enter valid book ID"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	err = services.DeleteBookService(int(ID))
	if err != nil {
		finalMsg.Msg = "No book found"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	finalMsg.Msg = "Delete request successful"
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(finalMsg)
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.Write(res)
}

func GetBookAnyway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	finalMsg := types.CustomBookResponse{}
	// Getting query data
	bookId := r.URL.Query().Get("bookId")
	authorId := r.URL.Query().Get("authorId")

	tempBook, errBook := strconv.ParseInt(bookId, 0, 0)
	tempAuthor, errAuthor := strconv.ParseInt(authorId, 0, 0)

	books := services.GetBookService(int(tempBook), int(tempAuthor))

	if bookId != "" && errBook != nil {
		finalMsg.Msg = "Invalid book Id"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	} else if authorId != "" && errAuthor != nil {
		finalMsg.Msg = "Invalid author Id"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

	res, err := json.Marshal(books)
	if len(books) == 0 {
		finalMsg.Msg = "No books found"
		res, err = json.Marshal(finalMsg)
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
		finalMsg.Msg = "Invalid book Id"
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(finalMsg)
		w.Write(res)
		return
	}
	books := services.GetBookService(int(ID), 0)
	if len(books) == 0 {
		finalMsg.Msg = "No books found"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	} else {
		book := models.Book{
			ID:          books[0].ID,
			Name:        books[0].Name,
			Publication: books[0].Publication,
			AuthorID:    books[0].AuthorID,
		}
		if updateBook.Name != "" {
			book.Name = updateBook.Name
		}
		if updateBook.Publication != "" {
			book.Publication = updateBook.Publication
		}
		if updateBook.AuthorID > 0 {
			book.AuthorID = updateBook.AuthorID
		}
		if err = book.Validate(); err != nil {
			finalMsg.Msg = err.Error()
			finalMsg.ErrorMsg = err
			w.WriteHeader(http.StatusBadRequest)
		} else {
			finalMsg.Content, finalMsg.ErrorMsg = services.UpdateBookService(book)
			w.WriteHeader(http.StatusCreated)
		}
	}
	if finalMsg.ErrorMsg != nil {
		finalMsg.Msg = "Update request failed"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	} else {
		finalMsg.Msg = "Update request successful"
	}

	res, err := json.Marshal(finalMsg)
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.Write(res)
}
