package sqlite

import (
	"database/sql"
	"posgo/internal/domain"
)

type PermissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) GetAllUserPermissions(userID int) (*domain.Permission, error) {
	query :=
		`	SELECT * FROM permissions WHERE id = ?
		`
	p := &domain.Permission{}

	err := r.db.QueryRow(query, userID).Scan(
		&p.ID,
		&p.UserID,
		&p.CreateProduct,
		&p.DeleteProduct,
		&p.ModifiedUnitCostPriceProduct,
		&p.ModifiedUnitSellPriceProduct,
		&p.ModifiedStockProduct,
		&p.ModifiedDiscount,
		&p.ModifiedProductName,
		&p.ModifiedAvailableDiscount,
		&p.ModifiedDepartment,
		&p.ModifiedUserName,
		&p.ModifiedUserLastName,
		&p.ModifiedUserPhoneNumber,
		&p.ModifiedUserPermissions,
		&p.ModifiedUserPassword,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil

}

func (r *PermissionRepository) UpdatePermission(permission *domain.Permission) error {
	query := `
        UPDATE permissions
        SET create_product = ?, 
            delete_product = ?, 
            modified_unit_cost_price_product = ?, 
            modified_unit_sell_price_product = ?, 
            modified_stock_product = ?, 
            modified_discount = ?, 
            modified_name = ?, 
            modified_available_discount = ?, 
            modified_department = ?, 
            modified_user_name = ?, 
            modified_user_last_name = ?, 
            modified_user_phone_number = ?, 
            modified_user_permissions = ?, 
            modified_user_password = ?
        WHERE user_id = ?
    `

	_, err := r.db.Exec(query,
		permission.CreateProduct,
		permission.DeleteProduct,
		permission.ModifiedUnitCostPriceProduct,
		permission.ModifiedUnitSellPriceProduct,
		permission.ModifiedStockProduct,
		permission.ModifiedDiscount,
		permission.ModifiedProductName,
		permission.ModifiedAvailableDiscount,
		permission.ModifiedDepartment,
		permission.ModifiedUserName,
		permission.ModifiedUserLastName,
		permission.ModifiedUserPhoneNumber,
		permission.ModifiedUserPermissions,
		permission.ModifiedUserPassword,
		permission.UserID,
	)

	return err
}
