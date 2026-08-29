package domain

import "time"

// ready
type ProductRepository interface {
	GetProductByCode(code string) (*Product, error)
	GetProductByName(name string) (*Product, error)
	GetProductLowStock() ([]*Product, error)
	GetProductbyDepartment(department string) ([]*Product, error)

	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	RemoveProduct(code string) error
}

type PermissionRepository interface {
	GetAllUserPermissions(UserId int) (*Permission, error)
	UpdatePermission(permission *Permission) error
}

// pending
type DepartmentRepository interface {
	GetDepartmentByName(name string) (*Department, error)
	GetAllDepartments() ([]*Department, error)
	CreateDepartment(department *Department) error
	UpdateDepartment(department *Department) error
	RemoveDepartment(name string) error
}

type TicketRepository interface {
	CreateTicket(ticket *Ticket) error
	GetTicketByID(ticketID int) (*Ticket, error)
	GetAllCurrentDay(date time.Time) ([]*Ticket, error)
	GetAllCurrentMonth(date time.Time) ([]*Ticket, error)
	GetAllCurrentSession(date time.Time) ([]*Ticket, error)
	UpdateTicket(ticket *Ticket) error
	DeleteTicket(ticketID int) error
}

type TicketItemRepository interface {
	GetAllByTicketID(ticketID int) ([]*TicketItem, error)
	CreateTicketItem(ticketItem *TicketItem) error
	UpdateTicketItem(ticketItem *TicketItem) error
	RemoveTicketItem(ticketID int, productCode string) error
}

// pending
type SessionRepository interface {
	CreateSession(session *Session) error
	CloseSession(token string) error
	GetActiveSessionByToken(token string) (*Session, error)
}
