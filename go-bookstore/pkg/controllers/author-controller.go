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
	finalMsg.Content = (*types.ResponseAuthor)(&author)
	err := author.Validate()

	// Validation failure check
	if err != nil {
		finalMsg.Msg = "Create request failed"
		finalMsg.ErrorMsg = err
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	finalMsg.Content, finalMsg.ErrorMsg = services.CreateAuthorService(author)
	// Database failure check
	if finalMsg.ErrorMsg != nil {
		finalMsg.Msg = "Create request failed"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	finalMsg.Msg = "Create request successful"
	res, err := json.Marshal(finalMsg)
	// Marshalling failure check
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.Write(res)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	finalMsg := types.CustomOnlyResponse{}

	// Getting query data
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)

	// Invalid Author check
	if err != nil || ID < 1 {
		finalMsg.Msg = "Enter valid Author ID"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	// Database failure check
	err = services.DeleteAuthorService(int(ID))
	if err != nil {
		finalMsg.Msg = "No author found"
		res, _ := json.Marshal(finalMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	finalMsg.Msg = "Delete request successful"

	// Marshalling failure check
	res, err := json.Marshal(finalMsg)
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	// Getting query data
	authorId := r.URL.Query().Get("authorId")

	tempAuthor, errAuthor := strconv.ParseInt(authorId, 0, 0)
	authors := services.GetAuthorService(int(tempAuthor))
	res, err := json.Marshal(authors)

	if len(authors) == 0 {
		res, err = json.Marshal("Msg : No authors found")
	}

	// Marshalling failure check
	if err != nil {
		res, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	} else if authorId != "" && errAuthor != nil {
		// invalid Author check
		res, _ = json.Marshal("Msg : Invalid author Id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}
