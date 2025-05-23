package postgres

import (
	"context"
	"employee-service/internal/domain/models"
	"employee-service/internal/storage"
	"errors"
	"github.com/jackc/pgx/v5"
	"time"
)

func (s *Storage) SaveEmployee(ctx context.Context, employee models.Employee) (int64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := s.db.Begin(ctxWithTimeout)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var passportId, departmentId, employeeId int64
	if err = tx.QueryRow(ctx, "INSERT INTO passports (type ,number) VALUES ($1, $2) RETURNING id", employee.Passport.Type, employee.Passport.Number).Scan(&passportId); err != nil {
		return 0, err
	}

	if err = tx.QueryRow(ctx, "SELECT id FROM departments where name = $1", employee.Department.Name).Scan(&departmentId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, storage.ErrDepartmentNotFound
		}
		return 0, err
	}

	if err = tx.QueryRow(ctx, "INSERT INTO employees (name ,surname, phone, company_id, passport_id, department_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", employee.Name, employee.Surname, employee.Phone, employee.CompanyId, passportId, departmentId).Scan(&employeeId); err != nil {
		return 0, err
	}

	return employeeId, nil
}
