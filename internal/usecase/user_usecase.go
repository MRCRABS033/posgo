package usecase

import (
	"errors"
	"posgo/internal/domain"
	"strings"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetUserByID(userID int) (*domain.User, error) {
	if userID <= 0 {
		return nil, errors.New("El id no puede venir vacio")
	}

	user, err := uc.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario no encontrado.")
	}
	return user, nil
}

func (uc *UserUseCase) GetUserByName(name string) (*domain.User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("")
	}

	user, err := uc.repo.GetUserByName(name)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario no encontrado.")
	}

	return user, nil
}

func (uc *UserUseCase) CreateUser(user *domain.User) (*domain.User, error) {
	if err := user.Validation(); err != nil {
		return nil, err
	}
	user, err := uc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUser(user *domain.User) (*domain.User, error) {
	if err := user.ValidationForUpdate(); err != nil {
		return nil, err
	}

	existUser, err := uc.repo.GetUserByID(user.ID)

	if err != nil {
		return nil, err
	}

	if existUser == nil {
		return nil, errors.New("Usuario no encontrado.")
	}

	updateUser, err := uc.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return updateUser, nil
}
