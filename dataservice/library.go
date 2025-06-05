package dataservice

import (
	"database/sql"
	"go-projetcs/model"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func CreateBook(db *sql.DB, book model.Book) error {
	//insert the book record into the "books" table
	query := "INSERT INTO books (title, author, year) VALUES (?, ?, ?)"
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err

	}
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

func UpdateBook(db *sql.DB, book model.Book) error {

	// Update the book record in the "books" table
	query := "UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?"
	_, err := db.Exec(query, book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(db *sql.DB, book model.Book) error {

	// Delete the book record from the "books" table
	query := "DELETE FROM books WHERE id = ?"
	_, err := db.Exec(query, book.ID)
	if err != nil {
		return err
	}
	return nil
}
