package containers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeServer(router *mux.Router) {
	fmt.Println("Server starting.......")
	err := http.ListenAndServe("localhost:9010", router)
	if err != nil {
		panic("Server lost")
	}
}
