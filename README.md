# Bookstore REST API
This is a simple REST API written in Go with MySQL.

## API

| Endpoint             | Response          | Method |
|----------------------|-------------------|--------|
| /api/v1/authors      | The authors       | GET    |
| /api/v1/authors/{id} | A single author   | GET    |
| /api/v1/authors      | Creates an author | POST   |
| /api/v1/authors/{id} | Updates an author | PUT    |
| /api/v1/authors/{id} | Deletes an author | DELETE |
| /api/v1/books        | The books         | GET    |
| /api/v1/books/{id}   | A single book     | GET    |
| /api/v1/books        | Creates a book    | POST   |
| /api/v1/books/{id}   | Updates a book    | PUT    |
| /api/v1/books/{id}   | Deletes a book    | DELETE |

## Building from source
1. Create a datbase with the name `bookstore`.
2. Load the mysqldump you find [here](./db/bookstore.sql).
3. ```sh
    # clone the project
    git clone https://github.com/0l1v3rr/bookstore-api.git

    # cd into it
    cd bookstore-api

    # run the main with the following command:
    go run main.go
   ```