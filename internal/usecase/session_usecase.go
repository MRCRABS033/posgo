package usecase

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"posgo/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type SessionUseCase struct {
	sessionRepo domain.SessionRepository
	userRepo    domain.UserRepository
}

func NewSessionUseCase(sr domain.SessionRepository, ur domain.UserRepository) *SessionUseCase {
	return &SessionUseCase{
		sessionRepo: sr,
		userRepo:    ur,
	}
}

func generateSecureUUIDs() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "fallback-token-secure"
	}
	return hex.EncodeToString(bytes)
}

// 1. Iniciar sesión (Login / CreateSession)
func (uc *SessionUseCase) Login(username, password string) (*domain.Session, error) {

	user, err := uc.userRepo.GetUserByName(username)
	if err != nil {
		return nil, errors.New("usuario no encontrado desde la base de datos")
	}

	if user == nil {
		return nil, errors.New("usuario nulo")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("usuario o contrasena incorrectos")
	}

	new_uuid := generateSecureUUIDs()

	session := &domain.Session{
		UserID:   user.ID,
		UUIDs:    new_uuid,
		LoginAt:  time.Now(),
		IsActive: true,
	}

	err = uc.sessionRepo.CreateSession(session)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (uc *SessionUseCase) Logout(UUID string) error {
	if strings.TrimSpace(UUID) == "" {
		return errors.New("el token de sesión no puede estar vacío.")
	}

	session, err := uc.sessionRepo.GetActiveSessionByToken(UUID)
	if err != nil {
		return err
	}
	if session == nil {
		return errors.New("no se encontró una sesión activa con el UUID proporcionado.")
	}

	now := time.Now()
	session.LogoutAt = &now
	session.IsActive = false

	if err := session.ValidateForLogout(); err != nil {
		return err
	}

	return uc.sessionRepo.CloseSession(UUID)
}

// 3. Obtener sesión activa por token
func (uc *SessionUseCase) GetActiveSessionByToken(token string) (*domain.Session, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("el token de búsqueda no puede estar vacío.")
	}

	session, err := uc.sessionRepo.GetActiveSessionByToken(token)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errors.New("sesión no encontrada o inactiva.")
	}

	return session, nil
}
