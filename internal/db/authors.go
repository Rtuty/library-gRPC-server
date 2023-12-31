package db

import (
	"context"

	"modules/internal/models"
)

// authorMethodsHandler паттерн простая фабрика, получает указание операции и структуру автора, после чего производит манипуляцию в базе данных
func (db *DataBase) AuthorMethodsHandler(ctx context.Context, operation string, author models.Author) error {
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

			return duplicateError
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
func (db *DataBase) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	r, err := db.client.QueryContext(ctx, "select id, name, country from authors")
	if err != nil {
		return nil, rowsQueryError
	}

	authors, err := scanRows(r, "author")
	if err != nil {
		return nil, scanError
	}

	result, err := convertEntity[models.Author](authors)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAuthorsByBookName возвращает список авторов по наименованию книги
func (db *DataBase) GetAuthorsByBookName(ctx context.Context, title string) ([]models.Author, error) {
	q := `select a.id, a.name, a.country 
			 from authors a 
			inner join books b on b.author_id = a.id and b.title = ?`

	r, err := db.client.QueryContext(ctx, q, title)
	if err != nil {
		return nil, rowsQueryError
	}

	authors, err := scanRows(r, "author")
	if err != nil {
		return nil, scanError
	}

	result, err := convertEntity[models.Author](authors)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAuthorById получает автора по id
func (db *DataBase) GetAuthorById(ctx context.Context, id int64) (models.Author, error) {
	var a models.Author

	if err := db.client.QueryRowContext(ctx, "select a.id, a.name, a.country from authors a where a.id = ?", id).Scan(&a.ID, &a.Name, &a.Country); err != nil {
		return models.Author{}, rowsQueryError
	}

	return a, nil
}
