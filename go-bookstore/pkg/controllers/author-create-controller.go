package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deadking/go-bookstore/pkg/models"
	"github.com/deadking/go-bookstore/pkg/types"
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
