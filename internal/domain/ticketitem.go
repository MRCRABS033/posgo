package domain

type TicketItem struct {
	Session  Session
	Product  Product
	Quantity float64
	Subtotal float64
}
