package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/types"
	"github.com/gorilla/mux"
)

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	finalMsg := types.CustomAuthorDeleteResponse{}

	// Getting query data
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	finalMsg.Content, finalMsg.Books, finalMsg.Msg = AuthorInt.Delete(int(ID))
	res, err := json.Marshal(finalMsg)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
