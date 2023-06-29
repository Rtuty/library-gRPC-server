package db

import (
	"context"
	"modules/internal/models"
)

// POST/PUT/DELETE Books methods
func (db *dataBase) AddNewBook(ctx context.Context, book models.Book) error {
	return nil
}
func (db *dataBase) UpdateBook(ctx context.Context, id string, book models.Book) error {
	return nil
}
func (db *dataBase) DeleteBook(ctx context.Context, id string) error {
	return nil
}

// POST/PUT/DELETE Authors methods
func (db *dataBase) AddNewAuthor(ctx context.Context, book models.Book) error {
	return nil
}

func (db *dataBase) UpdateAuthor(ctx context.Context, id string, book models.Author) error {
	return nil
}
func (db *dataBase) DeleteAuthor(ctx context.Context, id string) error {
	return nil
}

// GET Books methods
func (db *dataBase) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return []models.Book{}, nil
}
func (db *dataBase) GetBookById(ctx context.Context, id string) (models.Book, error) {
	return models.Book{}, nil
}
func (db *dataBase) GetBooksByAuthorName(ctx context.Context, name string) ([]models.Book, error) {
	return []models.Book{}, nil
}

// GET Authors methods
func (db *dataBase) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	return []models.Author{}, nil
}
func (db *dataBase) GetAuthorById(ctx context.Context, id string) (models.Author, error) {
	return models.Author{}, nil
}
func (db *dataBase) GetAuthorsByBookName(ctx context.Context, name string) ([]models.Author, error) {
	return []models.Author{}, nil
}
