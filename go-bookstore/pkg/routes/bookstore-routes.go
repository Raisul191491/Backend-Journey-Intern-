package routes

import (
	"github.com/deadking/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisteredBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBookAnyway).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.GetBookAnyway).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
