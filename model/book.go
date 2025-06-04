package model

type Book struct {
	Title  string
	Author string
	Year   int
	ID     int `json:"id,omitempty"` // ID is optional for creation, but required for updates
}
