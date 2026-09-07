package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) CreateSession(session *domain.Session) error {
	query := `
        INSERT INTO sessions (user_id, uuid, login_at, logout_at, is_active)
        VALUES (?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query,
		session.UserID,
		session.UUIDs,
		session.LoginAt,
		session.LogoutAt,
		session.IsActive,
	)
	return err
}

func (r *SessionRepository) CloseSession(token string) error {
	query := `
        UPDATE sessions 
        SET logout_at = CURRENT_TIMESTAMP, is_active = 0 
        WHERE token = ?
    `
	result, err := r.db.Exec(query, token)
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

func (r *SessionRepository) GetActiveSessionByToken(token string) (*domain.Session, error) {
	query := `
        SELECT id, user_id, token, login_at, logout_at, is_active
        FROM sessions
        WHERE token = ? AND is_active = 1
    `

	s := &domain.Session{}
	err := r.db.QueryRow(query, token).Scan(
		&s.ID,
		&s.UserID,
		&s.UUIDs,
		&s.LoginAt,
		&s.LogoutAt,
		&s.IsActive,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return s, nil
}
