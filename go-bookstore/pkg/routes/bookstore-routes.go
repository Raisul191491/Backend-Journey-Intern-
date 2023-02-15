package routes

import (
	"github.com/deadking/go-bookstore/pkg/controllers"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

var RegisteredBookStoreRoutes = func(router *mux.Router) {
	// Database Migration
	//InitializeDatabse()

	// Book routes
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/author", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/book", controllers.GetBookAnyway).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	// Author routes
}

func InitializeDatabse() {
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
	// db.Migrator().DropTable(&models.Book{}, &models.Author{})
	// db.Migrator().CreateTable(&models.Author{})
	// db.Migrator().CreateTable(&models.Book{})
}
