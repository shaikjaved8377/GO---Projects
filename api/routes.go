package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	// Register your routes here
	// e.g., http.HandleFunc("/example", exampleHandler)
	http.HandleFunc("/create", Createhandler(db))
	http.HandleFunc("/books", GetBookhandler(db))
	http.HandleFunc("/update", UpdateBookHandler(db)) // <-- Add this line
	http.HandleFunc("/delete", DeleteBookHandler(db)) // <-- Add this line
}
