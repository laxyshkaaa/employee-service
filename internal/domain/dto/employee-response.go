package dto

import "employee-service/internal/domain/models"

type EmpResponse struct {
	Message    string `json:"message"`
	EmployeeId int64  `json:"employee_id"`
}

type EmpsResponse struct {
	Message   string            `json:"message"`
	Employees []models.Employee `json:"employees"`
}
