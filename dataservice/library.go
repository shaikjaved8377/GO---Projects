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
	// Decode the JSON request body into the Book struct
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}
	//insert the book record into the "books" table
	query := "INSERT INTO books (title, author, year) VALUES (?, ?, ?)"
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err

	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}

func GetBooks(db *sql.DB) ([]model.Book, error) {
	rows, err := db.Query("SELECT title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Title, &book.Author, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func UpdateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	// Decode the JSON request body into the Book struct
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}
	// Update the book record in the "books" table
	query := "UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?"
	_, err := db.Exec(query, book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
	return nil
}

func DeleteBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	// Decode the JSON request body into the Book struct
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}
	// Delete the book record from the "books" table
	query := "DELETE FROM books WHERE id = ?"
	_, err := db.Exec(query, book.ID)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent) // No content for successful deletion
	return nil
}
