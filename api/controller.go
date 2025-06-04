package api

import (
	"database/sql"
	"net/http"
)

// Createhandler handles post requests to create a new book
func Createhandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := CreateBookLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func GetBookhandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := GetBookLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
