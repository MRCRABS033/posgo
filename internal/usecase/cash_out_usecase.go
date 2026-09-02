package usecase

import (
	"errors"
	"fmt"
	"posgo/internal/domain"
)

type CashOutUseCase struct {
	repo domain.CashOutRepository
}

func NewCashOutUseCase(repo domain.CashOutRepository) *CashOutUseCase {
	return &CashOutUseCase{repo: repo}
}

func (uc *CashOutUseCase) GetCashOutByID(cashOutID int) (*domain.CashOut, error) {
	if cashOutID <= 0 {
		return nil, errors.New("Ingrese un id valido.")
	}
	cashOut, err := uc.repo.GetCashOutByID(cashOutID)
	if err != nil {
		return nil, err
	}
	if cashOut == nil {
		return nil, fmt.Errorf("No se encontro el movimiento de efectivo con el id: %d", cashOutID)
	}
	return cashOut, nil
}

func (uc *CashOutUseCase) GetAllCashOutByUserID(userID int) ([]*domain.CashOut, error) {
	if userID <= 0 {
		return nil, errors.New("Ingrese un id de usuario valido.")
	}
	cashOutsUser, err := uc.repo.GetAllCashOutByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(cashOutsUser) <= 0 {
		return nil, fmt.Errorf("No se encontraron salidas de efectivo de este usuario.")
	}
	return cashOutsUser, nil

}

func (uc *CashOutUseCase) CreateCashOut(cashOut *domain.CashOut) (*domain.CashOut, error) {
	if err := cashOut.Validate(); err != nil {
		return nil, err
	}
	cashOut, err := uc.repo.CreateCashOut(cashOut)
	if err != nil {
		return nil, err
	}

	return cashOut, nil
}

func (uc *CashOutUseCase) UpdateCashOut(cashOut *domain.CashOut) (*domain.CashOut, error) {
	if err := cashOut.Validate(); err != nil {
		return nil, err
	}
	_, err := uc.repo.GetCashOutByID(cashOut.ID)
	if err != nil {
		return nil, errors.New("No se encontro una salida de efectivo con ese ID")
	}
	return uc.repo.UpdateCashOut(cashOut)

}
