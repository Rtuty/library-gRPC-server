package db

import (
	"context"

	"modules/internal/models"
)

// authorMethodsHandler паттерн простая фабрика, получает указание операции и структуру автора, после чего производит манипуляцию в базе данных
func (db *dataBase) authorMethodsHandler(ctx context.Context, operation string, author models.Author) error {
	var query string
	var args []any

	switch operation {
	case "create": // добавляет нового автора в базу данных, предварительно проверяя, существует ли создаваемый объект в бд
		var exists bool

		db.client.QueryRowContext(ctx, "select exists(select 1 from authors where name = ?)", author.Name).Scan(&exists)
		if exists {
			return dublicateError
		}

		query = "insert into authors (name, country) values (?, ?)"
		args = []interface{}{author.Name, author.Country}

	case "update": //обновляет поля в таблице authors в базе данных
		query = "update authors set name = ?, country = ? where id = ?"
		args = []interface{}{author.Name, author.Country, author.ID}

	case "delete": // удаляет сущность author из базы данных
		query = "delete from authors where id = ?"
		args = []interface{}{author.ID}

	}

	if err := executeQuery(ctx, db.client, query, args); err != nil {
		return err
	}

	return nil
}

// GetAllAuthors возвращает всех авторов из базы данных
func (db *dataBase) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	q := "select id, name, country from authors"

	r, err := db.client.QueryContext(ctx, q)
	if err != nil {
		return nil, rowsQueryError
	}

	authors, err := scanAuthorRows(r)
	if err != nil {
		return nil, scanError
	}

	return authors, nil
}

func (db *dataBase) GetAuthorsByBookName(ctx context.Context, title string) ([]models.Author, error) {
	q := `select a.id, a.name, a.country 
			 from authors a 
			inner join books b on b.author_id = a.id and b.title = ?`

	r, err := db.client.QueryContext(ctx, q, title)
	if err != nil {
		return nil, rowsQueryError
	}

	authors, err := scanAuthorRows(r)
	return []models.Author{}, nil

	if err != nil {
		return nil, scanError
	}

	return authors, nil
}

func (db *dataBase) GetAuthorById(ctx context.Context, id string) (models.Author, error) {
	return models.Author{}, nil
}
