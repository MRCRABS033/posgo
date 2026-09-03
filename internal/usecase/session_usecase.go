package usecase

import (
	"errors"
	"strings"
	"time"

	"posgo/internal/domain"
)

type SessionUseCase struct {
	repo domain.SessionRepository
}

func NewSessionUseCase(repo domain.SessionRepository) *SessionUseCase {
	return &SessionUseCase{repo: repo}
}

// 1. Iniciar sesión (Login / CreateSession)
func (uc *SessionUseCase) Login(session *domain.Session) error {
	if session == nil {
		return errors.New("los datos de la sesión son obligatorios.")
	}

	// Aplicamos la regla de negocio del dominio para el login
	if err := session.ValidateForLogin(); err != nil {
		return err
	}

	// Guardamos la sesión mediante el repositorio
	return uc.repo.CreateSession(session)
}

// 2. Cerrar sesión (Logout / CloseSession)
func (uc *SessionUseCase) Logout(token string) error {
	if strings.TrimSpace(token) == "" {
		return errors.New("el token de sesión no puede estar vacío.")
	}

	// Opcional: Podrías buscar primero la sesión activa para validar sus campos antes de cerrarla,
	// o delegar directamente la actualización en el repositorio mediante el token.
	session, err := uc.repo.GetActiveSessionByToken(token)
	if err != nil {
		return err
	}
	if session == nil {
		return errors.New("no se encontró una sesión activa con el token proporcionado.")
	}

	// Actualizamos los campos necesarios para el cierre en el dominio
	now := time.Now()
	session.LogoutAt = &now
	session.IsActive = false

	// Validamos el cierre usando el método de tu entidad
	if err := session.ValidateForLogout(); err != nil {
		return err
	}

	// Cerramos la sesión en la base de datos
	return uc.repo.CloseSession(token)
}

// 3. Obtener sesión activa por token
func (uc *SessionUseCase) GetActiveSessionByToken(token string) (*domain.Session, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("el token de búsqueda no puede estar vacío.")
	}

	session, err := uc.repo.GetActiveSessionByToken(token)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errors.New("sesión no encontrada o inactiva.")
	}

	return session, nil
}
