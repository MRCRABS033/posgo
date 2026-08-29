package usecase

import (
	"errors"
	"posgo/internal/domain"
)

type TicketItemUseCase struct {
	repo domain.TicketItemRepository
}

func NewTicketItemUseCase(repo domain.TicketItemRepository) *TicketItemUseCase {
	return &TicketItemUseCase{repo: repo}
}

func (uc *TicketItemUseCase) GetAllByTicketID(ticketID int) ([]*domain.TicketItem, error) {
	if ticketID <= 0 {
		return nil, errors.New("Por favor ingrese un ID de ticket válido.")
	}

	ticketItems, err := uc.repo.GetAllByTicketID(ticketID)
	if err != nil {
		return nil, err
	}
	if len(ticketItems) == 0 {
		return nil, errors.New("No se encontraron items para el ticket especificado.")
	}

	return ticketItems, nil
}

func (uc *TicketItemUseCase) CreateTicketItem(ticketItem *domain.TicketItem) error {
	if ticketItem == nil {
		return errors.New("Por favor ingrese un item de ticket válido.")
	}

	if err := uc.repo.CreateTicketItem(ticketItem); err != nil {
		return err
	}
	return nil
}

func (uc *TicketItemUseCase) UpdateTicketItem(ticketItem *domain.TicketItem) error {

	if ticketItem == nil {
		return errors.New("Por favor ingrese un item de ticket válido.")
	}

	if err := uc.repo.UpdateTicketItem(ticketItem); err != nil {
		return err
	}
	return nil
}

func (uc *TicketItemUseCase) RemoveTicketItem(ticketID int, productCode string) error {
	if ticketID <= 0 || productCode == "" {
		return errors.New("Por favor ingrese un item de ticket válido.")
	}

	if err := uc.repo.RemoveTicketItem(ticketID, productCode); err != nil {
		return err
	}
	return nil
}
