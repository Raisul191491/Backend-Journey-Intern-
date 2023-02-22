package main

import (
	"github.com/deadking/go-bookstore/pkg/containers"
)

func main() {

	// Initialize Database
	DB := containers.InitializeDatabse()

	// Initialize Interface
	containers.Initializeinterfaces(DB)

	// initalize Routing
	router := containers.InitalizeRouter()

	// Initialize server
	containers.InitializeServer(router)
}
