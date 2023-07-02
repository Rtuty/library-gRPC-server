package db

import (
	"context"

	"modules/internal/models"
)

// authorMethodsHandler паттерн простая фабрика, получает указание операции и структуру автора, после чего производит манипуляцию в базе данных
func (db *dataBase) AuthorMethodsHandler(ctx context.Context, operation string, author models.Author) error {
	t, err := db.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	switch operation {
	case "create": // добавляет нового автора в базу данных, предварительно проверяя, существует ли создаваемый объект в бд
		var exists bool

		t.QueryRowContext(ctx, "select exists(select 1 from authors where name = ?)", author.Name).Scan(&exists)
		if exists {
			if err = t.Rollback(); err != nil {
				return rollbackError
			}

			return dublicateError
		}

		if _, err = t.ExecContext(ctx, "insert into authors (name, country) values (?, ?)", author.Name, author.Country); err != nil {
			return execError
		}

	case "update": //обновляет поля в таблице authors в базе данных TODO Не работает
		if _, err = t.ExecContext(ctx, "update authors set name = ?, country = ? where id = ?", author.Name, author.Country, author.ID); err != nil {
			return execError
		}

	case "delete": // удаляет сущность author из базы данных
		if _, err = t.ExecContext(ctx, "delete from authors where id = ?", author.ID); err != nil {
			return execError
		}
	}

	if err = t.Commit(); err != nil {
		return commitError
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
	if err != nil {
		return nil, scanError
	}

	return authors, nil
}

func (db *dataBase) GetAuthorById(ctx context.Context, id string) (models.Author, error) {
	return models.Author{}, nil
}
