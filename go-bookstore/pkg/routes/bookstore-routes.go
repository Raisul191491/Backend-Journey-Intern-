package routes

import (
	"github.com/deadking/go-bookstore/pkg/controllers"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/repositories"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

var RegisteredBookStoreRoutes = func(router *mux.Router, db *gorm.DB) {

	// InitializeDatabse()
	Ibookcrud := repositories.BookDbInstance(db)
	Iauthorcrud := repositories.AuthorDbInstance(db)
	controllers.BookInterfaceInstance(Ibookcrud)
	controllers.AuthorInterfaceInstance(Iauthorcrud)

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

func InitializeDatabse() {
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
	// db.Migrator().DropTable(&models.Book{}, &models.Author{})
	// db.Migrator().CreateTable(&models.Author{})
	// db.Migrator().CreateTable(&models.Book{})
}
