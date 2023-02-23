package containers

import (
	"net/http"

	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/repositories"
	"github.com/deadking/go-bookstore/pkg/routes"
	"github.com/deadking/go-bookstore/pkg/services"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	config.Connect()
	db := config.GetDB()
	db.Migrator().AutoMigrate(&models.Author{})
	db.Migrator().AutoMigrate(&models.Book{})
	// db.Migrator().DropTable(&models.Book{}, &models.Author{})
	// db.Migrator().CreateTable(&models.Author{})
	// db.Migrator().CreateTable(&models.Book{})

	Ibookcrud := repositories.BookDbInstance(db)
	Iauthorcrud := repositories.AuthorDbInstance(db)
	services.BookInterfaceInstance(Ibookcrud)
	services.AuthorInterfaceInstance(Iauthorcrud)

	r := mux.NewRouter()
	routes.RegisteredBookStoreRoutes(r)
	http.Handle("/", r)
	return r
}
