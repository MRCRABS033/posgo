package usecase

import (
	"errors"
	"posgo/internal/domain"
	"strings"
)

type TicketItemUseCase struct {
	repo domain.TicketItemRepository
}

func NewTicketItemUseCase(repo domain.TicketItemRepository) *TicketItemUseCase {
	return &TicketItemUseCase{repo: repo}
}

func (uc *TicketItemUseCase) ModifyQuantity(ticketID int, productCode string, amount float64, action func(*domain.TicketItem) error) error {
	if ticketID <= 0 {
		return errors.New("Por favor ingrese un ID de ticket válido.")
	}
	if strings.TrimSpace(productCode) == "" {
		return errors.New("El código de producto no puede estar vacío.")
	}
	if amount <= 0 {
		return errors.New("La cantidad a ingresar no puede ser igual o menor a 0.")
	}
	item, err := uc.repo.GetByTicketIDAndProductCode(ticketID, productCode)
	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("El item no se encuentra en el ticket.")
	}

	if err := action(item); err != nil {
		return err
	}

	return uc.repo.UpdateTicketItem(item)

}

func (uc *TicketItemUseCase) AddQuantity(ticketID int, productCode string, amount float64) error {
	return uc.ModifyQuantity(ticketID, productCode, amount, func(item *domain.TicketItem) error {
		return item.AddQuantity(amount)
	})
}

func (uc *TicketItemUseCase) ReduceQuantity(ticketID int, productCode string, amount float64) error {
	return uc.ModifyQuantity(ticketID, productCode, amount, func(item *domain.TicketItem) error {
		return item.ReduceQuantity(amount)
	})
}

func (uc *TicketItemUseCase) GetByTicketIDAndProductCode(ticketID int, productCode string) (*domain.TicketItem, error) {
	if ticketID <= 0 {
		return nil, errors.New("Por favor ingrese un ID valido")
	}

	if strings.TrimSpace(productCode) == "" {
		return nil, errors.New("El codigo del producto no puede estar vacio.")
	}

	ticketItem, err := uc.repo.GetByTicketIDAndProductCode(ticketID, productCode)
	if err != nil {
		return nil, err
	}
	if ticketItem == nil {
		return nil, errors.New("No se encontro el producto.")
	}

	return ticketItem, nil
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
