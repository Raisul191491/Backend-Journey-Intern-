package containers

import (
	"net/http"

	"github.com/deadking/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func InitalizeRouter() *mux.Router {
	r := mux.NewRouter()
	routes.RegisteredBookStoreRoutes(r)
	http.Handle("/", r)
	return r
}
