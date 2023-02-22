package routes

import (
	"github.com/deadking/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

var RegisteredBookStoreRoutes = func(router *mux.Router) {

	// Book routes
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBookAnyway).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	// Author routes
	router.HandleFunc("/author", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/author", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.DeleteAuthor).Methods("DELETE")
}
