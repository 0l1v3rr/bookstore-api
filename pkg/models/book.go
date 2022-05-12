package models

import (
	"database/sql"
	"strconv"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      Author `json:"author"`
}

func BookRowMapper(rows *sql.Rows) ([]Book, error) {
	var books []Book

	columns, err := rows.Columns()
	if err != nil {
		return []Book{}, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var book Book
		id, _ := strconv.Atoi(string(values[0]))
		authorId, _ := strconv.Atoi(string(values[3]))
		book.ID = id
		book.Title = string(values[1])
		book.Description = string(values[2])
		book.Author = GetAuthorById(authorId)[0]

		books = append(books, book)
	}

	return books, nil
}

func GetAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		return []Book{}, err
	}

	books, err := BookRowMapper(rows)
	if err != nil {
		return []Book{}, err
	}

	return books, nil
}

func GetBookById(id int) (Book, error) {
	rows, err := db.Query("SELECT * FROM books WHERE id = ?;", id)
	if err != nil {
		return Book{}, err
	}

	book, err := BookRowMapper(rows)
	if err != nil {
		return Book{}, err
	}

	return book[0], nil
}
