package sqlite

import (
	"database/sql"
	_ "embed"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//go:embed sql/init.sql
var initSQL string

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func InitDB(db *sql.DB) error {
	_, err := db.Exec(initSQL)
	if err != nil {
		return err
	}

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		if err := seedDefaultUser(db); err != nil {
			return err
		}
	}
	return nil
}

func seedDefaultUser(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	plainPassword := "admin123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userQuery := `
		INSERT INTO users (name, last_name, password, phone_number)
		VALUES (?, ?, ?, ?)
		`
	result, err := tx.Exec(userQuery, "Admin", "Sistema", string(hashedPassword), "000 000 0000")

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	permQuery := `
		INSERT INTO permissions (
			user_id, create_product, delete_product, modified_unit_cost_price_product,
			modified_unit_sell_price_product, modified_stock_product, modified_discount,
			modified_product_name, modified_available_discount, modified_department,
			modified_user_name, modified_user_last_name, modified_user_phone_number,
			modified_user_permissions, modified_user_password
		) VALUES (?, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	`
	_, err = tx.Exec(permQuery, userID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Println("¡Usuario administrador por defecto creado con éxito (Contraseña protegida con bcrypt)!")
	return nil
}
