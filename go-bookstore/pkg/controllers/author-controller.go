package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
	"github.com/gorilla/mux"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	author := models.Author{}
	finalMsg := types.CustomAuthorResponse{}

	// Create row
	json.NewDecoder(r.Body).Decode(&author)
	finalMsg.Content, finalMsg.Msg = AuthorInt.Create(author)

	res, err := json.Marshal(finalMsg)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	finalMsg := types.CustomDeleteResponse{}

	// Getting query data
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	finalMsg.Msg = AuthorInt.Delete(int(ID))
	res, err := json.Marshal(finalMsg)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {

	// Getting query data
	tempAuthorID := r.URL.Query().Get("authorId")

	authorId := tempAuthorID
	if tempAuthorID == "" {
		authorId = defaultAuthorID
	}

	tempAuthor, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Enter a valid author ID :", err)
	}

	authors := AuthorInt.Get(int(tempAuthor))

	res, err := json.Marshal(authors)
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}