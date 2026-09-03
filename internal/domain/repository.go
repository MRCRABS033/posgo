package domain

import "time"

// ready
type CashInRepository interface {
	GetCashInByID(cashInID int) (*CashIn, error)
	GetAllCashInByUserID(userID int) ([]*CashIn, error)
	CreateCashIn(cashIn *CashIn) (*CashIn, error)
	UpdateCashIn(cashIn *CashIn) (*CashIn, error)
}

// ready
type CashOutRepository interface {
	GetCashOutByID(cashOutID int) (*CashOut, error)
	GetAllCashOutByUserID(userID int) ([]*CashOut, error)
	CreateCashOut(cashOut *CashOut) (*CashOut, error)
	UpdateCashOut(cashOut *CashOut) (*CashOut, error)
}

// ready
type DepartmentRepository interface {
	GetDepartmentByName(name string) (*Department, error)
	GetAllDepartments() ([]*Department, error)
	CreateDepartment(department *Department) error
	UpdateDepartment(department *Department) error
	RemoveDepartment(name string) error
}

// ready
type PermissionRepository interface {
	GetAllUserPermissions(UserId int) (*Permission, error)
	UpdatePermission(permission *Permission) error
}

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

// ready
type TicketItemRepository interface {
	AddQuantity(amount float64) error
	ReduceQuantity(amount float64) error
	GetByTicketIDAndProductCode(ticketID int, productCode string) (*TicketItem, error)
	GetAllByTicketID(ticketID int) ([]*TicketItem, error)
	CreateTicketItem(ticketItem *TicketItem) error
	UpdateTicketItem(ticketItem *TicketItem) error
	RemoveTicketItem(ticketID int, productCode string) error
}

// ready
type TicketRepository interface {
	CreateTicket(ticket *Ticket) (*Ticket, error)
	GetTicketByID(ticketID int) (*Ticket, error)
	GetAllCurrentDay(date time.Time) ([]*Ticket, error)
	GetAllCurrentMonth(date time.Time) ([]*Ticket, error)
	GetAllCurrentSession(date time.Time) ([]*Ticket, error)
	UpdateTicket(ticket *Ticket) error
	DeleteTicket(ticketID int) error
}

// pending
type SessionRepository interface {
	CreateSession(session *Session) error
	CloseSession(token string) error
	GetActiveSessionByToken(token string) (*Session, error)
}

type UserRepository interface {
	GetUserByID(userID int) (*User, error)
	GetUserByName(name string) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	ModifiedPermissionsUser(user *User) (*User, error)
	RemoveUser(user *User) (*User, error)
}
