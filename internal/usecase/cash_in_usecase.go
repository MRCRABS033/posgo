package usecase

import (
	"errors"
	"fmt"
	"posgo/internal/domain"
)

type CashInUseCase struct {
	repo domain.CashInRepository
}

func NewCashInUseCase(repo domain.CashInRepository) *CashInUseCase {
	return &CashInUseCase{repo: repo}
}

func (uc *CashInUseCase) GetCashInByID(cashInID int) (*domain.CashIn, error) {
	if cashInID <= 0 {
		return nil, errors.New("Por favor ingrese un ID de entrada de efectivo valido.")
	}

	cashIn, err := uc.repo.GetCashInByID(cashInID)
	if err != nil {
		return nil, err
	}

	if cashIn == nil {
		return nil, fmt.Errorf("No se encontro el movimiento de efectivo con el id: %d", cashInID)
	}

	return cashIn, nil
}

func (uc *CashInUseCase) GetAllCashInByUserID(userID int) ([]*domain.CashIn, error) {
	if userID <= 0 {
		return nil, errors.New("Por favor ingrese un ID de usuario valido.")
	}

	allCashInUser, err := uc.repo.GetAllCashInByUserID(userID)

	if err != nil {
		return nil, err
	}
	if len(allCashInUser) == 0 {
		return nil, errors.New("No se encontraron entradas de efectivo de este usuario.")
	}

	return allCashInUser, nil
}

func (uc *CashInUseCase) CreateCashIn(cashIn *domain.CashIn) (*domain.CashIn, error) {

	if err := cashIn.Validate(); err != nil {
		return nil, err
	}

	cashIn, err := uc.repo.CreateCashIn(cashIn)
	if err != nil {
		return nil, err
	}

	return cashIn, nil
}

func (uc *CashInUseCase) UpdateCashIn(cashIn *domain.CashIn) (*domain.CashIn, error) {
	if err := cashIn.Validate(); err != nil {
		return nil, err
	}
	_, err := uc.repo.GetCashInByID(cashIn.ID)
	if err != nil {
		return nil, errors.New("no se encontro una entrada de efectivo con ese ID")
	}
	return uc.repo.UpdateCashIn(cashIn)
}
