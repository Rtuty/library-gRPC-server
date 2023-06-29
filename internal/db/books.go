package db

import (
	"context"
	"modules/internal/models"
)

// POST/PUT/DELETE Books methods
func (db *dataBase) CreateNewBook(book models.Book) error {
	return nil
}
func (db *dataBase) UpdateBook(id string, book models.Book) error {
	return nil
}
func (db *dataBase) DeleteBook(id string) error {
	return nil
}

// GET Books methods
func (db *dataBase) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return []models.Book{}, nil
}
func (db *dataBase) GetBookById(id string) (models.Book, error) {
	return models.Book{}, nil
}
func (db *dataBase) GetBooksByAuthorName(name string) ([]models.Book, error) {
	return []models.Book{}, nil
}
