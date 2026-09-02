package domain

import (
	"errors"
	"strings"
)

type User struct {
	ID          int
	Name        string
	LastName    string
	Password    string
	PhoneNumber string
	permission  *Permission
}

type UserSession struct {
	ID     int64
	UserID int64
}

func (u *User) Validation() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("El nombre no puede estar vacio.")
	}
	if strings.TrimSpace(u.LastName) == "" {
		return errors.New("El apellido no puede quedar vacio")
	}
	if strings.TrimSpace(u.Password) == "" {
		return errors.New("La contrasena no puede estar vacia.")
	}
	if len(u.Password) < 6 {
		return errors.New("La contrasena no puede ser menor a 6 caracteres")
	}
	if u.permission == nil {
		return errors.New("Los permisos no pueden venir vacios.")
	}
	return nil
}

func (u *User) ValidationForUpdate() error {
	if u.ID <= 0 {
		return errors.New("El id de el usuario es invalido para actualizar.")
	}
	if err := u.Validation(); err != nil {
		return err
	}
	return nil
}
