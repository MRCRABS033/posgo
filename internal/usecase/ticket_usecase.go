package usecase

import (
	"errors"
	"posgo/internal/domain"
	"time"
)

type TicketUseCase struct {
	repo domain.TicketRepository
}

func NewTicketUseCase(repo domain.TicketRepository) *TicketUseCase {
	return &TicketUseCase{repo: repo}
}

func (uc *TicketUseCase) CreateTicket(ticket *domain.Ticket) error {
	if ticket == nil {
		return errors.New("Por favor ingrese un ticket válido.")
	}

	if err := uc.repo.CreateTicket(ticket); err != nil {
		return err
	}
	return nil
}

func (uc *TicketUseCase) GetTicketByID(ticketID int) (*domain.Ticket, error) {
	if ticketID <= 0 {
		return nil, errors.New("Por favor ingrese un ID de ticket válido.")
	}

	ticket, err := uc.repo.GetTicketByID(ticketID)

	if err != nil {
		return nil, err
	}
	if ticket == nil {
		return nil, errors.New("Ticket no encontrado")
	}

	return ticket, nil
}

func (uc *TicketUseCase) GetAllCurrentDay(date time.Time) ([]*domain.Ticket, error) {
	if date.IsZero() {
		return nil, errors.New("Por favor ingrese una fecha válida.")
	}

	allTickets, err := uc.repo.GetAllCurrentDay(date)

	if err != nil {
		return nil, err
	}

	if len(allTickets) == 0 {
		return nil, errors.New("No se encontraron tickets para el día actual.")
	}

	return allTickets, nil
}

func (uc *TicketUseCase) GetAllCurrentMonth(date time.Time) ([]*domain.Ticket, error) {
	if date.IsZero() {
		return nil, errors.New("Por favor ingrese una fecha válida.")
	}

	allTickets, err := uc.repo.GetAllCurrentMonth(date)

	if err != nil {
		return nil, err
	}

	if len(allTickets) == 0 {
		return nil, errors.New("No se encontraron tickets para el mes actual.")
	}

	return allTickets, nil
}

func (uc *TicketUseCase) GetAllCurrentSession(date time.Time) ([]*domain.Ticket, error) {
	if date.IsZero() {
		return nil, errors.New("Por favor ingrese una fecha válida.")
	}

	allTickets, err := uc.repo.GetAllCurrentSession(date)

	if err != nil {
		return nil, err
	}

	if len(allTickets) == 0 {
		return nil, errors.New("No se encontraron tickets para la sesión actual.")
	}

	return allTickets, nil
}

func (uc *TicketUseCase) UpdateTicket(ticket *domain.Ticket) error {
	if ticket == nil {
		return errors.New("Por favor ingrese un ticket válido.")
	}

	if err := uc.repo.UpdateTicket(ticket); err != nil {
		return err
	}
	return nil
}

func (uc *TicketUseCase) DeleteTicket(ticketID int) error {
	if ticketID <= 0 {
		return errors.New("Por favor ingrese un ID de ticket válido.")
	}

	if err := uc.repo.DeleteTicket(ticketID); err != nil {
		return err
	}
	return nil
}
