package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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