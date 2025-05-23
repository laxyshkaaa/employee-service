package postgres

import (
	"context"
	"employee-service/internal/domain/models"
	"time"
)

func (s *Storage) EmployeesByCompany(ctx context.Context, companyId int64) ([]models.Employee, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctxWithTimeout, `
		SELECT e.name, e.surname, e.phone, e.company_id,
		       d.name, d.phone,
		       p.type, p.number
		FROM employees e
		JOIN passports p ON e.passport_id = p.id
		JOIN departments d ON e.department_id = d.id
		WHERE e.company_id = $1
	`, companyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee

		err = rows.Scan(
			&emp.Name,
			&emp.Surname,
			&emp.Phone,
			&emp.CompanyId,
			&emp.Department.Name,
			&emp.Department.Phone,
			&emp.Passport.Type,
			&emp.Passport.Number,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	return employees, nil
}
