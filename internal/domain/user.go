package domain

type User struct {
	ID          int
	Name        string
	LastName    string
	Password    string
	PhoneNumber string
	permission  Permission
}

type UserSession struct {
	ID     int64
	UserID int64
}
