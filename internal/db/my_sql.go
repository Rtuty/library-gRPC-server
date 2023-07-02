package db

import (
	"context"

	"modules/internal/models"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Методы базы данных
type Repository interface {
	// operation (create, update, delete)
	BookMethodsHandler(ctx context.Context, operation string, book models.Book) error       // factory handler
	AuthorMethodsHandler(ctx context.Context, operation string, author models.Author) error //factory handler

	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookById(ctx context.Context, id string) (models.Book, error)
	GetBooksByAuthorId(ctx context.Context, id string) ([]models.Book, error)

	GetAllAuthors(ctx context.Context) ([]models.Author, error)
	GetAuthorById(ctx context.Context, id string) (models.Author, error)
	GetAuthorsByBookName(ctx context.Context, title string) ([]models.Author, error) // Возвращаем массив авторов, так как у одной книги их может быть несколько
}

type dataBase struct {
	client *sql.DB
}

func NewMySqlRepository(connection *sql.DB) Repository {
	return &dataBase{client: connection}
}
