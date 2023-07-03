package db

import (
	"database/sql"
	"errors"
	"modules/internal/models"
)

var (
	// local methods errors
	dublicateError = errors.New("the entity being created already exists in the database")
	// client methods errors
	scanError      = errors.New("error when scanning strings after query execution")
	rowsQueryError = errors.New("executed query returning strings failed with an error")
	// transaction errors
	beginError    = errors.New("failed to open transaction")
	execError     = errors.New("executing sql query error")
	commitError   = errors.New("failed to commit transaction")
	rollbackError = errors.New("failed to rollback transaction")
)

func scanAuthorRows(r *sql.Rows) ([]models.Author, error) {
	var authors []models.Author

	for r.Next() {
		var a models.Author

		if err := r.Scan(&a.ID, &a.Name, &a.Country); err != nil {
			return nil, scanError
		}

		authors = append(authors, a)
	}

	if err := r.Err(); err != nil {
		return nil, scanError
	}

	return authors, nil
}

func scanBooksRows(r *sql.Rows) ([]models.Book, error) {
	var books []models.Book

	for r.Next() {
		var b models.Book

		if err := r.Scan(&b.ID, &b.Title, &b.AuthorId, &b.Description); err != nil {
			return nil, scanError
		}

		books = append(books, b)
	}

	if err := r.Err(); err != nil {
		return nil, scanError
	}

	return books, nil
}
