package domain

import "time"

type ProductRepository interface {
	GetProductByCode(code string) (*Product, error)
	GetProductLowStock() ([]*Product, error)
	GetProductbyDepartment(department string) ([]*Product, error)

	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	RemoveProduct(code string) error
}

type DepartmentRepository interface {
	GetDepartmentByName(name string) (*Department, error)
	GetAllDepartments() ([]*Department, error)
	CreateDepartment(department *Department) error
	UpdateDepartment(department *Department) error
	RemoveDepartment(name string) error
}

type TicketRepository interface {
	GetAllCurrentDay() ([]*Department, error)
	GetAllCurrentMonth() ([]*Department, error)
	GetAllCurrentSession() ([]*Department, error)
}

type SessionRepository interface {
	CreateSession(session *Session) error
	CloseSession(token string) error
	GetActiveSessionByToken(token string) (*Session, error)
	GetByDate(date time.Time) ([]*Session, error)
}
