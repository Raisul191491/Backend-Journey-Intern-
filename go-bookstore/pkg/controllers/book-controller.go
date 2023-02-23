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
			Response(w, 400, "Create request failed, Author does not exist")
			return
		}
		finalMsg.Msg = "Create request successful"
	} else {
		Response(w, 400, finalMsg.ErrorMsg.Error())
	}
	Response(w, 200, finalMsg.Msg)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	// Getting query data
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil || ID < 1 {
		Response(w, 400, "Enter valid book ID")
		return
	}
	err = services.DeleteBookService(int(ID))
	if err != nil {
		Response(w, 400, "No book found")
		return
	}
	Response(w, 400, "Delete request successful")
}

func GetBookAnyway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	// Getting query data
	bookId := r.URL.Query().Get("bookId")

	authorId := r.URL.Query().Get("authorId")

	tempBook, errBook := strconv.ParseInt(bookId, 0, 0)
	tempAuthor, errAuthor := strconv.ParseInt(authorId, 0, 0)

	if bookId != "" && errBook != nil {
		Response(w, 400, "Enter valid book ID")
		return
	} else if authorId != "" && errAuthor != nil {
		Response(w, 400, "Enter valid author ID")
		return
	}
	books := services.GetBookService(int(tempBook), int(tempAuthor))

	res, _ := json.Marshal(books)
	if len(books) == 0 {
		Response(w, 200, "No books found")
	} else {
		Response(w, 200, "Books found")
		w.Write(res)
	}
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
		Response(w, 400, "Invalid book Id")
		return
	}
	books := services.GetBookService(int(ID), 0)
	if len(books) == 0 {
		Response(w, 200, "No books found")
		return
	} else {
		book := FormatStruct(updateBook, books)
		if err = book.Validate(); err != nil {
			Response(w, 400, finalMsg.ErrorMsg.Error())
		} else {
			finalMsg.Content, finalMsg.ErrorMsg = services.UpdateBookService(book)
		}
	}
	if finalMsg.ErrorMsg != nil {
		Response(w, 400, "Update request failed")
		return
	} else {
		Response(w, 201, "Update request successful")
	}
}

func FormatStruct(updateBook models.Book, books []types.ResponseBook) models.Book {
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
	return book
}
