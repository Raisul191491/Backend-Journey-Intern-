package controllers

import (
	"net/http"

	"github.com/deadking/go-bookstore/pkg/models"
)

var NewBook models.Book
var BasicID string = "0"

func GetBookAnyway(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Query().Get("bookId")
	// bookId := q
	// if q == "" {
	// 	bookId = BasicID
	// }
	// temp, err := strconv.ParseInt(bookId, 0, 0)
	// if err != nil {
	// 	fmt.Println("Enter a valid Book ID :", err)
	// }
	// bookDetails, _ := models.GetBookAnyway(temp)
	// res, _ := json.Marshal(bookDetails)
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// finalmsg := make(map[string]interface{})
	// createBook := &models.Book{}
	// utils.ParseBody(r, createBook)
	// b, msg := createBook.CreateBook()
	// finalmsg["Response"] = msg
	// finalmsg["Content"] = b
	// res, _ := json.Marshal(finalmsg)
	// w.WriteHeader(http.StatusCreated)
	// w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// finalmsg := make(map[string]interface{})
	// vars := mux.Vars(r)
	// bookId := vars["bookId"]
	// // fmt.Println(bookId)
	// ID, err := strconv.ParseInt(bookId, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing")
	// }
	// _, msg := models.DeleteBook(ID)
	// finalmsg["Response"] = msg
	// // finalmsg["Content"] = book
	// res, _ := json.Marshal(finalmsg)
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// finalmsg := make(map[string]interface{})
	// var updateBook = &models.Book{}
	// utils.ParseBody(r, updateBook)
	// vars := mux.Vars(r)
	// bookId := vars["bookId"]
	// ID, err := strconv.ParseInt(bookId, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing")
	// }
	// var res []byte
	// bookDetails, _ := models.GetBookAnyway(ID)
	// if len(bookDetails) > 0 {
	// 	if updateBook.Name != "" {
	// 		bookDetails[0].Name = updateBook.Name
	// 	}
	// 	if updateBook.Author != "" {
	// 		bookDetails[0].Author = updateBook.Author
	// 	}
	// 	if updateBook.Publication != "" {
	// 		bookDetails[0].Publication = updateBook.Publication
	// 	}
	// 	finalBook, msg := bookDetails[0].UpdateBook()
	// 	finalmsg["Response"] = msg
	// 	finalmsg["Content"] = finalBook
	// 	res, _ = json.Marshal(finalmsg)
	// } else {
	// 	res, _ = json.Marshal("Book not available")
	// }
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusAccepted)
	// w.Write(res)
}
