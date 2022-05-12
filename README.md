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