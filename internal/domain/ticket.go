package domain

import "time"

type Ticket struct {
	session   Session
	Id        int
	CreatedAt time.Time
	Items     []TicketItem
	Total     float64
}
