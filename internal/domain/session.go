package domain

import (
	"errors"
	"strings"
	"time"
)

type Session struct {
	ID       int
	UserID   int
	Token    string
	LoginAt  time.Time
	LogoutAt *time.Time
	IsActive bool
}

func (s *Session) Validate() error {
	if s == nil {
		return errors.New("datos de sesion invalidos.")
	}

	if s.UserID <= 0 {
		return errors.New("El ID del usuario no puede estar vacio")
	}
	if strings.TrimSpace(s.Token) == "" {
		return errors.New("El token no puede venir vacio.")
	}

	if s.LoginAt.IsZero() {
		return errors.New("La fecha de inicio de sesion no puede estar vacia.")
	}
	if !s.IsActive {
		return errors.New("Al iniciar sesion debe estar activo.")
	}
	return nil
}

func (s *Session) ValidateForLogin() error {

	if s == nil {
		return errors.New("datos de sesion invalidos para iniciar sesion.")
	}

	if s.UserID <= 0 {
		return errors.New("El id del usuario no puede estar vacio.")
	}

	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *Session) ValidateForLogout() error {

	if s == nil {
		return errors.New("datos de sesión inválidos.")
	}

	if s.ID <= 0 {
		return errors.New("El ID de la sesión no puede estar vacío.")
	}

	if s.LogoutAt == nil || s.LogoutAt.IsZero() {
		return errors.New("La fecha de cierre de sesion  no puede estar vacia.")
	}
	return nil
}
