package postgres

import (
	"context"

	"employee-service/internal/storage"
	"errors"
	"github.com/jackc/pgx/v5"
	"time"
)

func (s *Storage) DeleteEmployee(ctx context.Context, empId int) error {
	ctxWithTmt, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx, err := s.db.Begin(ctxWithTmt)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctxWithTmt)

	var passportId int64

	if err = tx.QueryRow(ctxWithTmt, "SELECT passport_id from employees where id = $1", empId).Scan(&passportId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.ErrEmployeesNotFound
		}
		return err
	}
	_, err = tx.Exec(ctxWithTmt, "DELETE FROM employees WHERE id = $1", empId)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctxWithTmt, "DELETE FROM passports WHERE id = $1", passportId)
	if err != nil {
		return err
	}
	if err = tx.Commit(ctxWithTmt); err != nil {
		return err
	}
	return nil

}
