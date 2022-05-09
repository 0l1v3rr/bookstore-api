package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/0l1v3rr/bookstore-api/pkg/models"
	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllAuthor()
	res, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}

	books := models.GetBookById(id)
	if len(books) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(books[0])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
