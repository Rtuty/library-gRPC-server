package db

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"

	"modules/internal/models"
)

// Методы базы данных
type Repository interface {
	CreateNewBook(book models.Book) error
	UpdateBook(id string, book models.Book) error
	DeleteBook(id string) error

	CreateNewAuthor(ctx context.Context, author models.Author) error
	UpdateAuthor(id string, book models.Author) error
	DeleteAuthor(id string) error

	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookById(id string) (models.Book, error)
	GetBooksByAuthorName(name string) ([]models.Book, error)

	GetAllAuthors() ([]models.Author, error)
	GetAuthorById(id string) (models.Author, error)
	GetAuthorsByBookName(name string) ([]models.Author, error) // Возвращаем массив авторов, так как у одной книги их может быть несколько
}

type dataBase struct {
	client *sql.DB
}

func NewMySqlConnection(connection *sql.DB) Repository {
	return &dataBase{client: connection}
}

var (
	dublicateEror = errors.New("the entity being created already exists in the database")
	// transaction errors
	beginError  = errors.New("failed to open transaction")
	commitError = errors.New("failed to commit transaction")
)
