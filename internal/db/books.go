package db

import (
	"context"
	"modules/internal/models"
)

// bookMethodsHandler паттерн простая фабрика, получает указание операции и структуру книги, после чего производит манипуляцию в базе данных
func (db *dataBase) bookMethodsHandler(ctx context.Context, operation string, book models.Book) error {
	var query string
	var args []any

	switch operation {
	case "create": // добавляет новую книгу в базу данных, предварительно проверяя, существует ли создаваемый объект в бд
		var exists bool

		db.client.QueryRowContext(ctx, "select exists(select 1 from books where id = ?)", book.ID).Scan(&exists)
		if exists {
			return dublicateError
		}

		query = "insert into books (title, author_id, description) values (?, ?, ?)"
		args = []any{book.Title, book.AuthorId, book.Description}

	case "update": // обновляет поля в таблице books в базе данных
		query = "update books set title = ?, author_id = ?, description = ? where id = ?"
		args = []any{book.Title, book.AuthorId, book.Description, book.ID}

	case "delete": // удаляет сущность book из базы данных
		query = "delete from authors where id = ?"
		args = []any{book.ID}

	}

	if err := executeQuery(ctx, db.client, query, args); err != nil {
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
