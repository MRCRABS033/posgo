package usecase

import (
	"errors"
	"posgo/internal/domain"
)

type PermissionUseCase struct {
	repo domain.PermissionRepository
}

func NewPermissionUseCase(repo domain.PermissionRepository) *PermissionUseCase {
	return &PermissionUseCase{repo: repo}
}

func (uc *PermissionUseCase) GetAllUserPermissions(UserId int) (*domain.Permission, error) {
	permissions, err := uc.repo.GetAllUserPermissions(UserId)

	if err != nil {
		return nil, err
	}
	if permissions == nil {
		return nil, errors.New("No se encontraron permisos para este usuario.")
	}

	return permissions, nil
}

func (uc *PermissionUseCase) UpdatePermission(permission *domain.Permission) error {

	if err := permission.Validate(); err != nil {
		return err
	}

	userPermissions, err := uc.repo.GetAllUserPermissions(permission.UserID)
	if err != nil {
		return err
	}

	if userPermissions == nil {
		return errors.New("No se encontraron permisos para este usuario.")
	}

	return uc.repo.UpdatePermission(permission)
}
