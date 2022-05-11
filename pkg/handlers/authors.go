package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/0l1v3rr/bookstore-api/pkg/models"
	"github.com/0l1v3rr/bookstore-api/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	utils.SetCors(w)

	books := models.GetAllAuthor()
	res, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	utils.SetCors(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	authors := models.GetAuthorById(id)
	if len(authors) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(authors[0])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	utils.SetCors(w)

	decoder := json.NewDecoder(r.Body)
	var author models.Author
	err := decoder.Decode(&author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = models.CreateAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
