package api

import (
	"database/sql"
	"go-projetcs/dataservice"
	"go-projetcs/model"
)

type IbizLogic interface {
	CreateBookLogic(book model.Book) error
	GetBookLogic() ([]model.Book, error)
	UpdateBookLogic(book model.Book) error
	DeleteBookLogic(book model.Book) error
}

type BizLogic struct {
	Db *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{Db: db}
}
func (bl *BizLogic) CreateBookLogic(book model.Book) error {
	return dataservice.CreateBook(bl.Db, book)

}

func (bl *BizLogic) GetBookLogic() ([]model.Book, error) {
	return dataservice.GetBooks(bl.Db)
}

func (bl *BizLogic) UpdateBookLogic(book model.Book) error {
	// Implement the logic to update a book in the database
	// This is a placeholder function; you need to implement the actual logic
	return dataservice.UpdateBook(bl.Db, book)
}

func (bl *BizLogic) DeleteBookLogic(book model.Book) error {
	// Implement the logic to delete a book from the database
	// This is a placeholder function; you need to implement the actual logic
	return dataservice.DeleteBook(bl.Db, book)
}
