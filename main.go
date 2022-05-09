package main

import (
	"log"
	"net/http"

	"github.com/0l1v3rr/bookstore-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/authors", handlers.GetAllAuthors).Methods("GET")
	r.HandleFunc("/api/v1/authors/{id}", handlers.GetAuthorById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
