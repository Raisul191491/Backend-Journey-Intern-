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
	err := author.Validate()
	if err == nil {
		finalMsg.Content, finalMsg.Msg = AuthorInt.Create(author)
	} else {
		finalMsg.Content, finalMsg.Msg = &models.Author{}, err.Error()
	}
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
	w.Header().Set("Content-Type", "pkglication/json")
	// Getting query data
	authorId := r.URL.Query().Get("authorId")

	tempAuthor, errAuthor := strconv.ParseInt(authorId, 0, 0)
	authors := AuthorInt.Get(int(tempAuthor))
	res, err := json.Marshal(authors)

	if len(authors) == 0 {
		res, err = json.Marshal("Msg : No authors found")
	}
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	} else if authorId != "" && errAuthor != nil {
		res, err = json.Marshal("Msg : Invalid author Id")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if err != nil {
		fmt.Println("Marshalling error", err.Error())
	}
	w.Write(res)
}
