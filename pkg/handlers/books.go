package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/0l1v3rr/bookstore-api/pkg/models"
	"github.com/0l1v3rr/bookstore-api/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	utils.SetCors(w)

	books, err := models.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	utils.SetCors(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
