package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductByCode(code string) (*domain.Product, error) {
	query := `
        SELECT code, name, stock, department, cost_price, sell_price
        FROM products
        WHERE code = ?
    `

	p := &domain.Product{}
	err := r.db.QueryRow(query, code).Scan(
		&p.Code, &p.Name, &p.Stock, &p.Department, &p.UnitCostPrice, &p.UnitSellPrice,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductRepository) GetProductByName(name string) (*domain.Product, error) {
	query := `
        SELECT code, name, stock, department, cost_price, sell_price
        FROM products
        WHERE name = ?
    `

	p := &domain.Product{}
	err := r.db.QueryRow(query, name).Scan(
		&p.Code, &p.Name, &p.Stock, &p.Department, &p.UnitCostPrice, &p.UnitSellPrice,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductRepository) GetProductLowStock() ([]*domain.Product, error) {
	query := `
        SELECT code, name, stock, department, cost_price, sell_price
        FROM products
        WHERE stock <= 5
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		p := &domain.Product{}
		if err := rows.Scan(
			&p.Code, &p.Name, &p.Stock, &p.Department, &p.UnitCostPrice, &p.UnitSellPrice,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductbyDepartment(department string) ([]*domain.Product, error) {
	query := `
        SELECT code, name, stock, department, cost_price, sell_price
        FROM products
        WHERE department = ?
    `

	rows, err := r.db.Query(query, department)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		p := &domain.Product{}
		if err := rows.Scan(
			&p.Code, &p.Name, &p.Stock, &p.Department, &p.UnitCostPrice, &p.UnitSellPrice,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	query := `
        INSERT INTO products (code, name, stock, department, cost_price, sell_price)
        VALUES (?, ?, ?, ?, ?, ?)
    `

	_, err := r.db.Exec(query,
		product.Code,
		product.Name,
		product.Stock,
		product.Department,
		product.UnitCostPrice,
		product.UnitSellPrice,
	)

	return err
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	query := `
        UPDATE products
        SET name = ?, stock = ?, department = ?, cost_price = ?, sell_price = ?
        WHERE code = ?
    `

	result, err := r.db.Exec(query,
		product.Name,
		product.Stock,
		product.Department,
		product.UnitCostPrice,
		product.UnitSellPrice,
		product.Code,
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

func (r *ProductRepository) RemoveProduct(code string) error {
	query := `
        DELETE FROM products
        WHERE code = ?
    `

	result, err := r.db.Exec(query, code)
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
