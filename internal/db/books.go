package db

import (
	"context"
	"modules/internal/models"
)

// POST/PUT/DELETE Books methods
func (db *dataBase) CreateNewBook(ctx context.Context, book models.Book) error {
	return nil
}
func (db *dataBase) UpdateBook(ctx context.Context, book models.Book) error {
	if err := executeQuery(ctx, db.client, "update books set title = ?, author_id = ?, description = ? where id = ?", []any{book.Title, book.AuthorId, book.Description, book.ID}); err != nil {
		return err
	}

	return nil
}
func (db *dataBase) DeleteBook(ctx context.Context, id string) error {
	if err := executeQuery(ctx, db.client, "delete from books where id = ?", []any{id}); err != nil {
		return err
	}

	return nil
}

// GET Books methods
func (db *dataBase) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return []models.Book{}, nil
}
func (db *dataBase) GetBookById(ctx context.Context, id string) (models.Book, error) {
	return models.Book{}, nil
}
func (db *dataBase) GetBooksByAuthorId(ctx context.Context, id string) ([]models.Book, error) {
	return []models.Book{}, nil
}
