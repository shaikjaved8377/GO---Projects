package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	// Register your routes here
	// e.g., http.HandleFunc("/example", exampleHandler)
	h := NewHandler(db)
	http.HandleFunc("/create", h.Createhandler(db))
	http.HandleFunc("/books", h.GetBookhandler(db))
	http.HandleFunc("/update", h.UpdateBookHandler(db)) // <-- Add this line
	http.HandleFunc("/delete", h.DeleteBookHandler(db)) // <-- Add this line
}
