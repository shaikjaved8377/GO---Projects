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
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(books)
}

func UpdateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Implement the logic to update a book in the database
	// This is a placeholder function; you need to implement the actual logic
	book := dataservice.UpdateBook(db, w, r)
	//if error != nil {
	//return error

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(book)
}

func DeleteBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Implement the logic to delete a book from the database
	// This is a placeholder function; you need to implement the actual logic
	if err := dataservice.DeleteBook(db, w, r); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent) // No content for successful deletion
	return nil
}
