package sqlite

import (
	"database/sql"
	"errors"
	"posgo/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(userID int) (*domain.User, error) {
	query := `SELECT * FROM users WHERE id = ?`

	u := &domain.User{
		Permission: &domain.Permission{},
	}

	err := r.db.QueryRow(query, userID).Scan(
		&u.ID,
		&u.Name,
		&u.LastName,
		&u.PhoneNumber,
		&u.Permission.ID,
		&u.Permission.UserID,
		&u.Permission.CreateProduct,
		&u.Permission.DeleteProduct,
		&u.Permission.ModifiedUnitCostPriceProduct,
		&u.Permission.ModifiedUnitSellPriceProduct,
		&u.Permission.ModifiedStockProduct,
		&u.Permission.ModifiedDiscount,
		&u.Permission.ModifiedProductName,
		&u.Permission.ModifiedAvailableDiscount,
		&u.Permission.ModifiedDepartment,
		&u.Permission.ModifiedUserName,
		&u.Permission.ModifiedUserLastName,
		&u.Permission.ModifiedUserPhoneNumber,
		&u.Permission.ModifiedUserPermissions,
		&u.Permission.ModifiedUserPassword,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetUserByName(name string) (*domain.User, error) {
	query := `
        SELECT 
            u.id, u.name, u.last_name, u.password, u.phone_number,
            p.id, p.user_id, p.create_product, p.delete_product,
            p.modified_unit_cost_price_product, p.modified_unit_sell_price_product,
            p.modified_stock_product, p.modified_discount, p.modified_product_name,
            p.modified_available_discount, p.modified_department, p.modified_user_name,
            p.modified_user_last_name, p.modified_user_phone_number,
            p.modified_user_permissions, p.modified_user_password
        FROM users u
        LEFT JOIN permissions p ON u.id = p.user_id
        WHERE u.name = ?
    `

	u := &domain.User{
		Permission: &domain.Permission{},
	}

	err := r.db.QueryRow(query, name).Scan(
		&u.ID,
		&u.Name,
		&u.LastName,
		&u.Password,
		&u.PhoneNumber,
		&u.Permission.ID,
		&u.Permission.UserID,
		&u.Permission.CreateProduct,
		&u.Permission.DeleteProduct,
		&u.Permission.ModifiedUnitCostPriceProduct,
		&u.Permission.ModifiedUnitSellPriceProduct,
		&u.Permission.ModifiedStockProduct,
		&u.Permission.ModifiedDiscount,
		&u.Permission.ModifiedProductName,
		&u.Permission.ModifiedAvailableDiscount,
		&u.Permission.ModifiedDepartment,
		&u.Permission.ModifiedUserName,
		&u.Permission.ModifiedUserLastName,
		&u.Permission.ModifiedUserPhoneNumber,
		&u.Permission.ModifiedUserPermissions,
		&u.Permission.ModifiedUserPassword,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	queryUser := `INSERT INTO users (name, last_name, phone_number) VALUES (?, ?, ?)`
	result, err := tx.Exec(queryUser, user.Name, user.LastName, user.PhoneNumber)
	if err != nil {
		return nil, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(userID)

	if user.Permission != nil {
		queryPerm := `
            INSERT INTO permissions (
                user_id, create_product, delete_product, 
                modified_unit_cost_price_product, modified_unit_sell_price_product, 
                modified_stock_product, modified_discount, modified_name, 
                modified_available_discount, modified_department, modified_user_name, 
                modified_user_last_name, modified_user_phone_number, 
                modified_user_permissions, modified_user_password
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        `
		_, err = tx.Exec(queryPerm,
			user.ID,
			user.Permission.CreateProduct,
			user.Permission.DeleteProduct,
			user.Permission.ModifiedUnitCostPriceProduct,
			user.Permission.ModifiedUnitSellPriceProduct,
			user.Permission.ModifiedStockProduct,
			user.Permission.ModifiedDiscount,
			user.Permission.ModifiedProductName,
			user.Permission.ModifiedAvailableDiscount,
			user.Permission.ModifiedDepartment,
			user.Permission.ModifiedUserName,
			user.Permission.ModifiedUserLastName,
			user.Permission.ModifiedUserPhoneNumber,
			user.Permission.ModifiedUserPermissions,
			user.Permission.ModifiedUserPassword,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	query := `
        UPDATE users 
        SET name = ?, last_name = ?, phone_number = ? 
        WHERE id = ?
    `
	result, err := r.db.Exec(query, user.Name, user.LastName, user.PhoneNumber, user.ID)
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

	return user, nil
}

func (r *UserRepository) ModifiedPermissionsUser(user *domain.User) (*domain.User, error) {
	if user.Permission == nil {
		return nil, errors.New("el usuario no cuenta con un objeto de permisos válido")
	}

	query := `
        UPDATE permissions 
        SET create_product = ?, delete_product = ?, 
            modified_unit_cost_price_product = ?, modified_unit_sell_price_product = ?, 
            modified_stock_product = ?, modified_discount = ?, modified_name = ?, 
            modified_available_discount = ?, modified_department = ?, modified_user_name = ?, 
            modified_user_last_name = ?, modified_user_phone_number = ?, 
            modified_user_permissions = ?, modified_user_password = ?
        WHERE user_id = ?
    `
	result, err := r.db.Exec(query,
		user.Permission.CreateProduct,
		user.Permission.DeleteProduct,
		user.Permission.ModifiedUnitCostPriceProduct,
		user.Permission.ModifiedUnitSellPriceProduct,
		user.Permission.ModifiedStockProduct,
		user.Permission.ModifiedDiscount,
		user.Permission.ModifiedProductName,
		user.Permission.ModifiedAvailableDiscount,
		user.Permission.ModifiedDepartment,
		user.Permission.ModifiedUserName,
		user.Permission.ModifiedUserLastName,
		user.Permission.ModifiedUserPhoneNumber,
		user.Permission.ModifiedUserPermissions,
		user.Permission.ModifiedUserPassword,
		user.ID, // Este corresponde al WHERE user_id = ?
	)
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

	return user, nil
}

func (r *UserRepository) RemoveUser(user *domain.User) (*domain.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM permissions WHERE user_id = ?`, user.ID)
	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(`DELETE FROM users WHERE id = ?`, user.ID)
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

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
