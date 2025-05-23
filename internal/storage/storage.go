package storage

import "errors"

var (
	ErrDepartmentNotFound = errors.New("department not found")
	ErrPassportNotFound   = errors.New("passport not found")
	ErrEmployeesNotFound  = errors.New("employees not found")
)
