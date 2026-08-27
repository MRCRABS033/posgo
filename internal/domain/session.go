package domain

import "time"

type Session struct {
	ID       int
	UserId   int
	Token    string
	LoginAt  time.Time
	LogoutAt *time.Time
	IsActive bool
}
