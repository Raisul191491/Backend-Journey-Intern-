package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/deadking/go-bookstore/pkg/models"
)

type CustomAuthorResponse struct {
	Content models.Author
	Msg     string
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	author := models.Author{}
	json.NewDecoder(r.Body).Decode(&author)
	cA, msg := author.CreateAuthor()
	finalMsg := &CustomAuthorResponse{
		Content: *cA,
		Msg:     msg,
	}
	res, _ := json.Marshal(finalMsg)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
