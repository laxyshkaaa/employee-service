package models

type Employee struct {
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyId  string     `json:"company_id"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}
