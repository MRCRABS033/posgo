package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
)

type TicketItemRepository struct {
	db *sql.DB
}

func NewTicketItemRepository(db *sql.DB) *TicketItemRepository {
	return &TicketItemRepository{db: db}
}

// 1. AddQuantity a nivel de repositorio (delegado o ejecutado directamente si aplica)
func (r *TicketItemRepository) AddQuantity(amount float64) error {
	// La validación y el cálculo de stock/subtotal principal residen en el Dominio.
	// Este método queda disponible para cumplir con la interfaz del contrato.
	return nil
}

// 2. ReduceQuantity a nivel de repositorio
func (r *TicketItemRepository) ReduceQuantity(amount float64) error {
	// Al igual que AddQuantity, la lógica de negocio pura vive en el Dominio
	// y se consolida en la base de datos a través de UpdateTicketItem.
	return nil
}

// 3. GetByTicketIDAndProductCode (Ya lo tenías definido)
func (r *TicketItemRepository) GetByTicketIDAndProductCode(ticketID int, productCode string) (*domain.TicketItem, error) {
	query := `
        SELECT 
            ti.ticket_id,
            ti.quantity,
            ti.subtotal,
            s.id,
            s.user_id,
            s.token,
            s.login_at,
            s.logout_at,
            s.is_active,
            p.id,
            p.code,
            p.name,
            p.unit_cost_price,
            p.unit_sell_price,
            p.discount,
            p.is_single_product,
            p.available_discount,
            p.stock,
            p.available,
            d.id,
            d.name
        FROM ticket_items ti
        JOIN sessions s ON ti.session_id = s.id
        JOIN products p ON ti.product_code = p.code
        LEFT JOIN departments d ON p.department_id = d.id
        WHERE ti.ticket_id = ? AND ti.product_code = ?
    `

	item := &domain.TicketItem{
		Session: &domain.Session{},
		Product: &domain.Product{
			Department: &domain.Department{},
		},
	}

	err := r.db.QueryRow(query, ticketID, productCode).Scan(
		&item.TicketID,
		&item.Quantity,
		&item.Subtotal,
		&item.Session.ID,
		&item.Session.UserID,
		&item.Session.Token,
		&item.Session.LoginAt,
		&item.Session.LogoutAt,
		&item.Session.IsActive,
		&item.Product.ID,
		&item.Product.Code,
		&item.Product.Name,
		&item.Product.UnitCostPrice,
		&item.Product.UnitSellPrice,
		&item.Product.Discount,
		&item.Product.IsSingleProduct,
		&item.Product.AvailableDiscount,
		&item.Product.Stock,
		&item.Product.Available,
		&item.Product.Department.ID,
		&item.Product.Department.Name,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return item, nil
}

// 4. GetAllByTicketID (Obtiene todos los ítems de un ticket con sus relaciones hidratadas)
func (r *TicketItemRepository) GetAllByTicketID(ticketID int) ([]*domain.TicketItem, error) {
	query := `
        SELECT 
            ti.ticket_id,
            ti.quantity,
            ti.subtotal,
            s.id,
            s.user_id,
            s.token,
            s.login_at,
            s.logout_at,
            s.is_active,
            p.id,
            p.code,
            p.name,
            p.unit_cost_price,
            p.unit_sell_price,
            p.discount,
            p.is_single_product,
            p.available_discount,
            p.stock,
            p.available,
            d.id,
            d.name
        FROM ticket_items ti
        JOIN sessions s ON ti.session_id = s.id
        JOIN products p ON ti.product_code = p.code
        LEFT JOIN departments d ON p.department_id = d.id
        WHERE ti.ticket_id = ?
    `

	rows, err := r.db.Query(query, ticketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.TicketItem

	for rows.Next() {
		item := &domain.TicketItem{
			Session: &domain.Session{},
			Product: &domain.Product{
				Department: &domain.Department{},
			},
		}

		if err := rows.Scan(
			&item.TicketID,
			&item.Quantity,
			&item.Subtotal,
			&item.Session.ID,
			&item.Session.UserID,
			&item.Session.Token,
			&item.Session.LoginAt,
			&item.Session.LogoutAt,
			&item.Session.IsActive,
			&item.Product.ID,
			&item.Product.Code,
			&item.Product.Name,
			&item.Product.UnitCostPrice,
			&item.Product.UnitSellPrice,
			&item.Product.Discount,
			&item.Product.IsSingleProduct,
			&item.Product.AvailableDiscount,
			&item.Product.Stock,
			&item.Product.Available,
			&item.Product.Department.ID,
			&item.Product.Department.Name,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, rows.Err()
}

// 5. CreateTicketItem (Inserta un nuevo ítem relacionando ticket, sesión y producto)
func (r *TicketItemRepository) CreateTicketItem(ticketItem *domain.TicketItem) error {
	query := `
        INSERT INTO ticket_items (ticket_id, session_id, product_code, quantity, subtotal)
        VALUES (?, ?, ?, ?, ?)
    `

	var sessionID int
	if ticketItem.Session != nil {
		sessionID = ticketItem.Session.ID
	}

	var productCode string
	if ticketItem.Product != nil {
		productCode = ticketItem.Product.Code
	}

	_, err := r.db.Exec(query,
		ticketItem.TicketID,
		sessionID,
		productCode,
		ticketItem.Quantity,
		ticketItem.Subtotal,
	)

	return err
}

// 6. UpdateTicketItem (Actualiza cantidad y subtotal tras sumar/restar)
func (r *TicketItemRepository) UpdateTicketItem(ticketItem *domain.TicketItem) error {
	query := `
        UPDATE ticket_items
        SET quantity = ?, subtotal = ?
        WHERE ticket_id = ? AND product_code = ?
    `

	var productCode string
	if ticketItem.Product != nil {
		productCode = ticketItem.Product.Code
	}

	result, err := r.db.Exec(query,
		ticketItem.Quantity,
		ticketItem.Subtotal,
		ticketItem.TicketID,
		productCode,
	)
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

// 7. RemoveTicketItem (Elimina un ítem específico del ticket)
func (r *TicketItemRepository) RemoveTicketItem(ticketID int, productCode string) error {
	query := `
        DELETE FROM ticket_items
        WHERE ticket_id = ? AND product_code = ?
    `

	result, err := r.db.Exec(query, ticketID, productCode)
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
