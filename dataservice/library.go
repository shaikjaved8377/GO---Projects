package dataservice

import (
	"database/sql"
	"encoding/json"
	"go-projetcs/model"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}
	query := "INSERT INTO books (title, author, year) VALUES (?, ?, ?)"

	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err

	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}
