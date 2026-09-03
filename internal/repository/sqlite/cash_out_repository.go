package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
)

type CashOutRepository struct {
	db *sql.DB
}

func NewCashOutRepository(db *sql.DB) *CashOutRepository {
	return &CashOutRepository{db: db}
}

func (r *CashOutRepository) GetCashOutByID(cashOutID int) (*domain.CashOut, error) {
	query :=
		`	SELECT id, user_id, session_id, concept, quantity
			FROM cash_outs
			WHERE id = ?
		`

	c := &domain.CashOut{}
	err := r.db.QueryRow(query, cashOutID).Scan(
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

func (r *CashOutRepository) GetAllCashOutByUserID(userID int) ([]*domain.CashOut, error) {
	query :=
		`	SELEECT id, user_id, session_id, concept, quantity
			FROM cash_outs
			WHERE id = ?
		`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cashOuts []*domain.CashOut

	for rows.Next() {
		c := &domain.CashOut{}

		if err := rows.Scan(
			&c.ID,
			&c.Quantity,
			&c.Concept,
			&c.SessionID,
			&c.UserID,
		); err != nil {
			return nil, err
		}

		cashOuts = append(cashOuts, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cashOuts, nil
}

func (r *CashOutRepository) CreateCashOut(c *domain.CashOut) (*domain.CashOut, error) {
	query :=
		`	INSERT INTO cash_outs (quantity, user_id, session_id, concept)
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

func (r *CashOutRepository) UpdateCashOut(c *domain.CashOut) (*domain.CashOut, error) {
	query :=
		`	UPDATE cash_outs
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
