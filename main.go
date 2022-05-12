package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0l1v3rr/bookstore-api/pkg/handlers"
	"github.com/gorilla/mux"
)

const PORT = 8080

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/authors", handlers.GetAllAuthors).Methods("GET")
	r.HandleFunc("/api/v1/authors/{id}", handlers.GetAuthorById).Methods("GET")
	r.HandleFunc("/api/v1/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/api/v1/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/api/v1/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), r))
}
