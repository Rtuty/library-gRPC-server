package db

import (
	"context"
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
	beginError  = errors.New("failed to open transaction")
	execError   = errors.New("executing sql query error")
	commitError = errors.New("failed to commit transaction")
)

// executeQuery открывает транзакцию и исполняет запрос
func executeQuery(ctx context.Context, c *sql.DB, q string, args []any) error {
	t, err := c.BeginTx(ctx, nil) //opts todo
	if err != nil {
		return beginError
	}

	if _, err = t.ExecContext(ctx, q, args); err != nil {
		return execError
	}

	if err = t.Commit(); err != nil {
		return commitError
	}

	return nil
}

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
