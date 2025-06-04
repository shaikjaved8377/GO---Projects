package api

import (
	"database/sql"
	"encoding/json"
	"go-projetcs/dataservice"
	"net/http"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Implement the logic to create a book in the database
	// This is a placeholder function; you need to implement the actual logi}
	return dataservice.CreateBook(db, w, r)
}

func GetBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Implement the logic to get a book from the database
	// This is a placeholder function; you need to implement the actual logic
	books, err := dataservice.GetBooks(db)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(books)
}
