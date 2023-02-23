package routes

import (
	"github.com/deadking/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

var RegisteredBookStoreRoutes = func(router *mux.Router) {

	// Book routes
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBookAnyway).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")

	// Author routes
	router.HandleFunc("/authors", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/authors/{authorId}", controllers.DeleteAuthor).Methods("DELETE")
}
