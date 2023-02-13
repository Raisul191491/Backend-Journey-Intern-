package routes

import (
	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/controllers"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var RegisteredBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBookAnyway).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.GetBookAnyway).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	InitializeDatabse()
}

func InitializeDatabse() {
	config.Connect()
	db = config.GetDB()
	db.DropTableIfExists(&models.Book{}, &models.Author{})
	db.AutoMigrate(&models.Book{}, &models.Author{})
	db.Model(&models.Book{}).AddForeignKey("author_id", "authors(id)", "CASCADE", "CASCADE")

}
