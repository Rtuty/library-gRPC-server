package db

import (
	"database/sql"
	"errors"
	"modules/internal/models"
)

var (
	// local methods errors
	dublicateError = errors.New("the entity being created already exists in the database")
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

// convertToAuthors Приводит сущности из пустого интерфейса к типу author
func convertToAuthors(entity []interface{}) ([]models.Author, error) {
	var authors []models.Author

	for _, e := range entity {
		if a, ok := e.(models.Author); ok {
			authors = append(authors, a)
		} else {
			return nil, convertError
		}
	}

	return authors, nil
}

// convertToBooks Приводит сущности из пустого интерфейса к типу book
func convertToBooks(entity []interface{}) ([]models.Book, error) {
	var books []models.Book

	for _, e := range entity {
		if b, ok := e.(models.Book); ok {
			books = append(books, b)
		} else {
			return nil, convertError
		}
	}

	return books, nil
}
