package db

import (
	"context"

	"modules/internal/models"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Методы базы данных
type Repository interface {
	CreateNewBook(ctx context.Context, book models.Book) error
	UpdateBook(ctx context.Context, book models.Book) error
	DeleteBook(ctx context.Context, id string) error

	CreateNewAuthor(ctx context.Context, author models.Author) error
	UpdateAuthor(ctx context.Context, author models.Author) error
	DeleteAuthor(ctx context.Context, id string) error

	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookById(ctx context.Context, id string) (models.Book, error)
	GetBooksByAuthorId(ctx context.Context, id string) ([]models.Book, error)

	GetAllAuthors(ctx context.Context) ([]models.Author, error)
	GetAuthorById(ctx context.Context, id string) (models.Author, error)
	GetAuthorsByBookName(ctx context.Context, name string) ([]models.Author, error) // Возвращаем массив авторов, так как у одной книги их может быть несколько
}

type dataBase struct {
	client *sql.DB
}

func NewMySqlConnection(connection *sql.DB) Repository {
	return &dataBase{client: connection}
}
