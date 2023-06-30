package db

import (
	"context"

	"modules/internal/models"
)

// CreateNewAuthor добавляет нового автора в базу данных, предварительно проверяя, существует ли создаваемый автор
func (db *dataBase) CreateNewAuthor(ctx context.Context, author models.Author) error {
	var exists bool

	db.client.QueryRowContext(ctx, "select exists(select 1 from authors where name = ?)", author.Name).Scan(&exists)

	if exists {
		return dublicateError
	}

	if err := executeQuery(ctx, db.client, "insert into authors (name, country) values (?, ?)", []any{author.Name, author.Country}); err != nil {
		return err
	}

	return nil
}

// UpdateAuthor обновляет поля в таблице author в базе данных
func (db *dataBase) UpdateAuthor(ctx context.Context, author models.Author) error {
	if err := executeQuery(ctx, db.client, "update authors set name = ?, country = ? where id = ?", []any{author.Name, author.Country, author.ID}); err != nil {
		return err
	}

	return nil
}

// DeleteAuthor удаляет сущность автор из базы данных
func (db *dataBase) DeleteAuthor(ctx context.Context, id string) error {
	if err := executeQuery(ctx, db.client, "delete from authors where id = ?", []any{id}); err != nil {
		return err
	}

	return nil
}

// GetAllAuthors возвращает всех авторов из базы данных
func (db *dataBase) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	return []models.Author{}, nil
}
func (db *dataBase) GetAuthorById(ctx context.Context, id string) (models.Author, error) {
	return models.Author{}, nil
}
func (db *dataBase) GetAuthorsByBookName(ctx context.Context, name string) ([]models.Author, error) {
	return []models.Author{}, nil
}
