package models

type Book struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      *Author `json:"author"`
}
