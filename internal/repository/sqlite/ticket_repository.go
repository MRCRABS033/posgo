package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
	"time"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) CreateTicket(ticket *domain.Ticket) (*domain.Ticket, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	sessionID := 0
	if ticket.Session != nil {
		sessionID = ticket.Session.ID
	}

	queryTicket := `
        INSERT INTO tickets (session_id, created_at, total, is_completed)
        VALUES (?, ?, ?, ?)
    `
	result, err := tx.Exec(queryTicket, sessionID, ticket.CreatedAt, ticket.Total, ticket.IsCompleted)
	if err != nil {
		return nil, err
	}

	ticketID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	ticket.Id = int(ticketID)

	queryItem := `
        INSERT INTO ticket_items (ticket_id, product_code, quantity, price)
        VALUES (?, ?, ?, ?)
    `
	for _, item := range ticket.Items {
		_, err := tx.Exec(queryItem, ticket.Id, item.Product.Code, item.Quantity, item.Product.UnitSellPrice)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepository) GetTicketByID(ticketID int) (*domain.Ticket, error) {
	queryTicket := `
        SELECT id, session_id, created_at, total, is_completed
        FROM tickets
        WHERE id = ?
    `

	t := &domain.Ticket{}
	var sessionID int

	err := r.db.QueryRow(queryTicket, ticketID).Scan(
		&t.Id,
		&sessionID,
		&t.CreatedAt,
		&t.Total,
		&t.IsCompleted,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	items, err := r.getItemsByTicketID(t.Id)
	if err != nil {
		return nil, err
	}
	t.Items = items

	return t, nil
}

func (r *TicketRepository) GetAllCurrentDay(date time.Time) ([]*domain.Ticket, error) {
	dateStr := date.Format("2006-01-02")
	query := `
        SELECT id, session_id, created_at, total, is_completed
        FROM tickets
        WHERE DATE(created_at) = ?
    `

	rows, err := r.db.Query(query, dateStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*domain.Ticket
	for rows.Next() {
		t := &domain.Ticket{}
		var sessionID int
		if err := rows.Scan(&t.Id, &sessionID, &t.CreatedAt, &t.Total, &t.IsCompleted); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, rows.Err()
}

func (r *TicketRepository) GetAllCurrentMonth(date time.Time) ([]*domain.Ticket, error) {
	monthStr := date.Format("2006-01")
	query := `
        SELECT id, session_id, created_at, total, is_completed
        FROM tickets
        WHERE strftime('%Y-%m', created_at) = ?
    `

	rows, err := r.db.Query(query, monthStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*domain.Ticket
	for rows.Next() {
		t := &domain.Ticket{}
		var sessionID int
		if err := rows.Scan(&t.Id, &sessionID, &t.CreatedAt, &t.Total, &t.IsCompleted); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, rows.Err()
}

func (r *TicketRepository) GetAllCurrentSession(date time.Time) ([]*domain.Ticket, error) {
	query := `
        SELECT id, session_id, created_at, total, is_completed
        FROM tickets
        WHERE DATE(created_at) = ?
    `
	rows, err := r.db.Query(query, date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*domain.Ticket
	for rows.Next() {
		t := &domain.Ticket{}
		var sessionID int
		if err := rows.Scan(&t.Id, &sessionID, &t.CreatedAt, &t.Total, &t.IsCompleted); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, rows.Err()
}

func (r *TicketRepository) UpdateTicket(ticket *domain.Ticket) error {
	query := `
        UPDATE tickets
        SET total = ?, is_completed = ?
        WHERE id = ?
    `
	result, err := r.db.Exec(query, ticket.Total, ticket.IsCompleted, ticket.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *TicketRepository) DeleteTicket(ticketID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM ticket_items WHERE ticket_id = ?`, ticketID)
	if err != nil {
		return err
	}

	result, err := tx.Exec(`DELETE FROM tickets WHERE id = ?`, ticketID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return tx.Commit()
}

func (r *TicketRepository) getItemsByTicketID(ticketID int) ([]*domain.TicketItem, error) {
	query := `
        SELECT product_code, quantity, price
        FROM ticket_items
        WHERE ticket_id = ?
    `
	rows, err := r.db.Query(query, ticketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.TicketItem
	for rows.Next() {
		item := &domain.TicketItem{}
		if err := rows.Scan(&item.Product.Code, &item.Quantity, &item.Product.UnitSellPrice); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}
