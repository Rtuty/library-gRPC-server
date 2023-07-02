package db

import (
	"context"
	"modules/internal/models"
)

// bookMethodsHandler паттерн простая фабрика, получает указание операции и структуру книги, после чего производит манипуляцию в базе данных
func (db *dataBase) BookMethodsHandler(ctx context.Context, operation string, book models.Book) error {
	t, err := db.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	switch operation {
	case "create": // добавляет новую книгу в базу данных, предварительно проверяя, существует ли создаваемый объект в бд
		var exists bool

		db.client.QueryRowContext(ctx, "select exists(select 1 from books where id = ?)", book.ID).Scan(&exists)
		if exists {
			return dublicateError
		}

		if _, err = t.ExecContext(ctx, "insert into books (title, author_id, description) values (?, ?, ?)", book.Title, book.AuthorId, book.Description); err != nil {
			return execError
		}

	case "update": // обновляет поля в таблице books в базе данных
		if _, err = t.ExecContext(ctx, "update books set title = ?, author_id = ?, description = ? where id = ?", book.Title, book.AuthorId, book.Description, book.ID); err != nil {
			return execError
		}

	case "delete": // удаляет сущность book из базы данных
		if _, err = t.ExecContext(ctx, "delete from authors where id = ?", book.ID); err != nil {
			return execError
		}
	}

	if err = t.Commit(); err != nil {
		return commitError
	}

	return nil
}

// GET Books methods
func (db *dataBase) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	q := "select id, title, author_id, description from books"

	r, err := db.client.QueryContext(ctx, q)
	if err != nil {
		return nil, rowsQueryError
	}

	books, err := scanBooksRows(r)
	if err != nil {
		return nil, scanError
	}

	return books, nil
}

func (db *dataBase) GetBookById(ctx context.Context, id string) (models.Book, error) {
	return models.Book{}, nil
}
func (db *dataBase) GetBooksByAuthorId(ctx context.Context, id string) ([]models.Book, error) {
	return []models.Book{}, nil
}
