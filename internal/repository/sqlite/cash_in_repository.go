package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
)

type CashInRepository struct {
	db *sql.DB
}

func NewCashInRepository(db *sql.DB) *CashInRepository {
	return &CashInRepository{db: db}
}

func (r *CashInRepository) GetCashInByID(cashInID int) (*domain.CashIn, error) {
	query :=
		`	SELECT id, user_id, session_id, concept, quantity
			FROM cash_ins
			WHERE id = ?
		`

	c := &domain.CashIn{}
	err := r.db.QueryRow(query, cashInID).Scan(
		&c.ID,
		&c.Quantity,
		&c.Concept,
		&c.SessionID,
		&c.UserID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CashInRepository) GetAllCashInByUserID(userID int) ([]*domain.CashIn, error) {
	query :=
		`	SELECT id, user_id, session_id, concept, quantity
			FROM cash_ins
			WHERE user_id = ?
		`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cashIns []*domain.CashIn

	for rows.Next() {
		c := &domain.CashIn{}

		if err := rows.Scan(
			&c.ID,
			&c.Quantity,
			&c.Concept,
			&c.SessionID,
			&c.UserID,
		); err != nil {
			return nil, err
		}
		cashIns = append(cashIns, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cashIns, nil
}

func (r *CashInRepository) CreateCashIn(c *domain.CashIn) (*domain.CashIn, error) {
	query :=
		`	INSERT INTO cash_ins (quantity, user_id, session_id, concept)
			VALUES (?, ?, ?, ?)
		`

	result, err := r.db.Exec(query, c.Quantity, c.UserID, c.SessionID, c.Concept)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	c.ID = int(id)
	return c, nil
}

func (r *CashInRepository) UpdateCashIn(c *domain.CashIn) (*domain.CashIn, error) {
	query :=
		`	UPDATE cash_ins
			SET quantity = ?, user_id = ?, session_id = ?, concept = ?
			WHERE id = ?
		`

	result, err := r.db.Exec(query, c.Quantity, c.UserID, c.SessionID, c.Concept, c.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	return c, nil
}
