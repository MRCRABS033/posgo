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

func (uc *TicketUseCase) CreateTicket(ticket *domain.Ticket) (*domain.Ticket, error) {

	if err := ticket.ValidateForCreation(); err != nil {
		return nil, err
	}

	newTicket, err := uc.repo.CreateTicket(ticket)

	if err != nil {
		return nil, err
	}
	if newTicket == nil {
		return nil, errors.New("no se pudo generar el ticket en la base de datos")
	}
	return newTicket, nil
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

	existTicket, err := uc.repo.GetTicketByID(ticket.Id)

	if err != nil {
		return err
	}

	if existTicket == nil {
		return errors.New("No se encontro un Ticket con este id")
	}

	existTicket.Total = ticket.Total
	existTicket.IsCompleted = ticket.IsCompleted

	if err := existTicket.ValidateForUpdate(); err != nil {
		return err
	}

	return uc.repo.UpdateTicket(existTicket)
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
