package main

import (
	"fmt"
	"net/http"

	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Database connect
	config.Connect()

	// Routing
	r := mux.NewRouter()
	routes.RegisteredBookStoreRoutes(r)
	http.Handle("/", r)

	// Initialize server
	fmt.Println("Server starting.......")
	err := http.ListenAndServe("localhost:9010", r)
	if err != nil {
		panic("Server lost")
	}
}
