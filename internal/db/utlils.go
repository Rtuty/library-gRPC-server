package db

import (
	"database/sql"
	"errors"
	"modules/internal/models"
)

var (
	// local methods errors
	duplicateError = errors.New("the entity being created already exists in the database")
	convertError   = errors.New("invalid item type in the array")
	// client methods errors
	scanError      = errors.New("error when scanning strings after query execution")
	rowsQueryError = errors.New("executed query returning strings failed with an error")
	// transaction errors
	beginError    = errors.New("failed to open transaction")
	execError     = errors.New("executing sql query error")
	commitError   = errors.New("failed to commit transaction")
	rollbackError = errors.New("failed to rollback transaction")
)

// scanRows сканирует строки логических сущностей, которые вернул sql запрос и возвращает их в виде пустого интерфейса
func scanRows(r *sql.Rows, modelType string) ([]interface{}, error) {
	var entities []interface{}

	for r.Next() {
		var entity interface{}

		switch modelType {
		case "author":
			var a models.Author
			if err := r.Scan(&a.ID, &a.Name, &a.Country); err != nil {
				return nil, scanError
			}
			entity = a

		case "book":
			var b models.Book
			if err := r.Scan(&b.ID, &b.Title, &b.AuthorId, &b.Description); err != nil {
				return nil, scanError
			}
			entity = b
		}

		entities = append(entities, entity)
	}

	if err := r.Err(); err != nil {
		return nil, scanError
	}

	return entities, nil
}

// convertEntity Конвертирует пустой интерфейс к массиву authors/books
func convertEntity[T models.Author | models.Book](entity []interface{}) ([]T, error) {
	var entities []T

	for _, e := range entity {
		if e1, ok := e.(T); ok {
			entities = append(entities, e1)
		} else {
			return make([]T, 0), convertError
		}
	}
	return entities, nil
}
