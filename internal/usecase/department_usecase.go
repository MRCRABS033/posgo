package usecase

import (
	"errors"
	"posgo/internal/domain"
	"strings"
)

type DepartmentUseCase struct {
	repo domain.DepartmentRepository
}

func NewDepartmentUseCase(repo domain.DepartmentRepository) *DepartmentUseCase {
	return &DepartmentUseCase{repo: repo}
}

func (uc *DepartmentUseCase) GetDepartmentByName(name string) (*domain.Department, error) {

	department, err := uc.repo.GetDepartmentByName(name)

	if err != nil {
		return nil, err
	}
	if department == nil {
		return nil, errors.New("No se encontró el departamento.")
	}

	return department, nil
}

func (uc *DepartmentUseCase) GetAllDepartments() ([]*domain.Department, error) {
	deparments, err := uc.repo.GetAllDepartments()
	if err != nil {
		return nil, err
	}
	if deparments == nil {
		return nil, errors.New("No se encontraron departamentos.")
	}

	return deparments, nil
}

func (uc *DepartmentUseCase) CreateDepartment(name string) (*domain.Department, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("Por favor escriba el nombre del departamento.")
	}

	existDepartment, err := uc.repo.GetDepartmentByName(name)
	if err != nil {
		return nil, err
	}
	if existDepartment != nil {
		return nil, errors.New("Ya existe un departamento con ese nombre.")
	}

	department := &domain.Department{
		Name: name,
	}

	err = uc.repo.CreateDepartment(department)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (uc *DepartmentUseCase) UpdateDepartment(name string) (*domain.Department, error) {
	if name == "" {
		return nil, errors.New("Por favor ingrese un nombre de departamento.")
	}

	existDepartment, err := uc.repo.GetDepartmentByName(name)
	if err != nil {
		return nil, err
	}
	if existDepartment == nil {
		return nil, errors.New("No se encontró el departamento.")
	}

	department := &domain.Department{
		Name: name,
	}
	return department, nil
}

func (uc *DepartmentUseCase) RemoveDepartment(name string) error {
	if name == "" {
		return errors.New("Por favor ingrese un nombre de departamento.")
	}

	existDepartment, err := uc.repo.GetDepartmentByName(name)
	if err != nil {
		return err
	}
	if existDepartment == nil {
		return errors.New("No se encontró el departamento.")
	}

	return uc.repo.RemoveDepartment(name)
}
