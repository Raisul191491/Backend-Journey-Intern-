package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/services"
	"github.com/deadking/go-bookstore/pkg/types"
	"github.com/gorilla/mux"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	author := models.Author{}
	finalMsg := types.CustomAuthorResponse{}

	// Create row
	json.NewDecoder(r.Body).Decode(&author)
	newAuthor := models.Author{
		AuthorName: author.AuthorName,
		Age:        author.Age,
	}
	finalMsg.Content = (*types.ResponseAuthor)(&newAuthor)
	err := author.Validate()

	// Validation failure check
	if err != nil {
		Response(w, 400, "Create request failed")
		return
	}

	finalMsg.Content, finalMsg.ErrorMsg = services.CreateAuthorService(author)
	// Database failure check
	if finalMsg.ErrorMsg != nil {
		Response(w, 400, "Create request failed")
		return
	}
	Response(w, 200, "Create request successful")
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")

	// Getting query data
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)

	// Invalid Author check
	if err != nil || ID < 1 {
		Response(w, 400, "Enter valid Author ID")
		return
	}

	// Database failure check
	err = services.DeleteAuthorService(int(ID))
	if err != nil {
		Response(w, 400, "No author found")
		return
	}
	// Marshalling failure check
	if err != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	Response(w, 406, "Delete request successful")
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	// Getting query data
	authorId := r.URL.Query().Get("authorId")

	tempAuthor, errAuthor := strconv.ParseInt(authorId, 0, 0)
	if authorId != "" && errAuthor != nil {
		// invalid Author check
		Response(w, 406, "Msg : Invalid author Id")
		return
	}
	authors := services.GetAuthorService(int(tempAuthor))
	res, err := json.Marshal(authors)

	if len(authors) == 0 {
		Response(w, 200, "Msg : No authors found")
		return
	}

	// Marshalling failure check
	if err != nil {
		Response(w, 406, "Marshalling error")
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}

func Response(w http.ResponseWriter, code int, msg string) {
	res, err := json.Marshal(msg)
	if err != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	w.WriteHeader(code)
	w.Write(res)
}
