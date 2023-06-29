package db

import (
	"context"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"modules/internal/models"
)

// Методы базы данных
type MySql interface {
	AddNewBook(ctx context.Context, book models.Book) error
	UpdateBook(ctx context.Context, id string, book models.Book) error
	DeleteBook(ctx context.Context, id string) error

	AddNewAuthor(ctx context.Context, book models.Book) error
	UpdateAuthor(ctx context.Context, id string, book models.Author) error
	DeleteAuthor(ctx context.Context, id string) error

	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookById(ctx context.Context, id string) (models.Book, error)
	GetBooksByAuthorName(ctx context.Context, name string) ([]models.Book, error)

	GetAllAuthors(ctx context.Context) ([]models.Author, error)
	GetAuthorById(ctx context.Context, id string) (models.Author, error)
	GetAuthorsByBookName(ctx context.Context, name string) ([]models.Author, error) // Возвращаем массив авторов, так как у одной книги их может быть несколько
}

type dataBase struct {
	db *sql.DB
}

func NewMySqlConnection(db *sql.DB) MySql {
	return &dataBase{db: db}
}
