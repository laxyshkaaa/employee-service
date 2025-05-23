package postgres

import (
	"context"
	"employee-service/internal/storage"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strings"
	"time"
)

func (s *Storage) UpdateEmployee(ctx context.Context, updateFields map[string]interface{}, empId int64) error {
	ctxWithTmt, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var passportId int64

	if passportType, ok := updateFields["passport_type"]; ok {
		tx, err := s.db.Begin(ctxWithTmt)
		if err != nil {
			return err
		}
		defer tx.Rollback(ctxWithTmt)

		err = tx.QueryRow(ctxWithTmt, "SELECT passport_id FROM employees WHERE id = $1", empId).Scan(&passportId)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return storage.ErrEmployeesNotFound
			}
			return err
		}

		_, err = tx.Exec(ctxWithTmt, "UPDATE passports SET type = $1 WHERE id = $2", passportType, passportId)
		if err != nil {
			return err
		}

		if err = tx.Commit(ctxWithTmt); err != nil {
			return err
		}
		delete(updateFields, "passport_type")
	}

	if passportNumber, ok := updateFields["passport_number"]; ok {
		tx, err := s.db.Begin(ctxWithTmt)
		if err != nil {
			return err
		}
		defer tx.Rollback(ctxWithTmt)

		err = tx.QueryRow(ctxWithTmt, "SELECT passport_id FROM employees WHERE id = $1", empId).Scan(&passportId)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return storage.ErrEmployeesNotFound
			}
			return err
		}

		_, err = tx.Exec(ctxWithTmt, "UPDATE passports SET number = $1 WHERE id = $2", passportNumber, passportId)
		if err != nil {
			return err
		}

		if err = tx.Commit(ctxWithTmt); err != nil {
			return err
		}
		delete(updateFields, "passport_number")
	}

	if len(updateFields) > 0 {
		query, args := PrepareQuerySafe(updateFields, empId)
		_, err := s.db.Exec(ctxWithTmt, query, args...)
		if err != nil {
			return err
		}
	}

	return nil
}

func PrepareQuerySafe(updateFields map[string]interface{}, empId int64) (string, []interface{}) {
	baseQuery := "UPDATE employees SET "
	parts := make([]string, 0, len(updateFields))
	args := make([]interface{}, 0, len(updateFields)+1)

	i := 1
	for k, v := range updateFields {
		parts = append(parts, fmt.Sprintf("%s = $%d", k, i))
		args = append(args, v)
		i++
	}

	query := baseQuery + strings.Join(parts, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, empId)

	return query, args
}
