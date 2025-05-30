package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	// Register your routes here
	// e.g., http.HandleFunc("/example", exampleHandler)
	http.HandleFunc("/create", Createhandler(db))
}
