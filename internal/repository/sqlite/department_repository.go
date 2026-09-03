package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
)

type DeparmentRepository struct {
	db *sql.DB
}

func NewDeparmentRepository(db *sql.DB) *DeparmentRepository {
	return &DeparmentRepository{db: db}
}

func (r *DeparmentRepository) GetDepartmentByName(name string) (*domain.Department, error) {
	query :=
		`	SELECT id, name
			FROM departments
			WHERE name = ?
		`
	d := &domain.Department{}
	err := r.db.QueryRow(query, name).Scan(
		&d.ID,
		&d.Name,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *DeparmentRepository) GetAllDepartments() ([]*domain.Department, error) {
	query :=
		`	SELECT id, name
			FROM departments
		`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var departments []*domain.Department

	for rows.Next() {
		d := &domain.Department{}

		if err := rows.Scan(
			&d.ID,
			&d.Name,
		); err != nil {
			return nil, err
		}

		departments = append(departments, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *DeparmentRepository) CreateDepartment(d *domain.Department) error {
	query :=
		`	INSERT INTO departments (name)
		VALUES (?)
	`

	result, err := r.db.Exec(query, d.Name)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	d.ID = int(id)
	return nil
}

func (r *DeparmentRepository) UpdateDepartment(d *domain.Department) error {
	query :=
		`	UPDATE departments
			SET name = ?
			WHERE id = ?
		`

	result, err := r.db.Exec(query, d.ID)
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

func (r *DeparmentRepository) RemoveDepartment(name string) error {
	query := `
        DELETE FROM departments
        WHERE name = ?
    `

	result, err := r.db.Exec(query, name)
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
