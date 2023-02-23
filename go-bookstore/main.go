package main

import (
	"fmt"
	"net/http"

	"github.com/deadking/go-bookstore/pkg/containers"
)

func main() {

	// Initialize Database
	router := containers.Init()

	fmt.Println("Server starting.......")
	err := http.ListenAndServe("localhost:9010", router)
	if err != nil {
		panic("Server lost")
	}
}
