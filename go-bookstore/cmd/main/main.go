package main

import (
	"fmt"
	"net/http"

	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteredBookStoreRoutes(r)
	http.Handle("/", r)
	models.InitializeDatabse()
	fmt.Println("Server starting.......")
	err := http.ListenAndServe("localhost:9010", r)
	if err != nil {
		panic("Jisjdisd")
	}
	// log.Fatal(http.ListenAndServe(":9010", r))
}
