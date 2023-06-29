package db

import (
	"context"
	"modules/internal/models"
)

// AddNewAuthor добавляет нового автора в базу данных, предварительно проверяя, существует ли создаваемый автор
func (db *dataBase) CreateNewAuthor(ctx context.Context, author models.Author) error {
	var exists bool

	db.client.QueryRowContext(ctx, "select exists(select 1 from authors where name = $1)", author.Name).Scan(&exists)

	if exists {
		return dublicateEror
	}

	t, err := db.client.BeginTx(ctx, nil)
	if err != nil {
		return beginError
	}

	// todo: придумать реализацию для книг автора
	if _, err := t.ExecContext(ctx, "insert into authors (name, books) values ($1, $2)", author.Name, author.Books); err != nil {
		return err
	}

	if err := t.Commit(); err != nil {
		return commitError
	}

	return nil
}

func (db *dataBase) UpdateAuthor(id string, book models.Author) error {
	return nil
}
func (db *dataBase) DeleteAuthor(id string) error {
	return nil
}

// GET Authors methods
func (db *dataBase) GetAllAuthors() ([]models.Author, error) {
	return []models.Author{}, nil
}
func (db *dataBase) GetAuthorById(id string) (models.Author, error) {
	return models.Author{}, nil
}
func (db *dataBase) GetAuthorsByBookName(name string) ([]models.Author, error) {
	return []models.Author{}, nil
}
