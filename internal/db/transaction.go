package db

import (
	"context"
	"database/sql"
	"errors"
)

var (
	dublicateError = errors.New("the entity being created already exists in the database")
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
