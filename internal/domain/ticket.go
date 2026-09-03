package domain

import (
	"errors"
	"time"
)

type Ticket struct {
	Session     *Session
	Id          int
	CreatedAt   time.Time
	Items       []*TicketItem
	Total       float64
	IsCompleted bool
}

func (t *Ticket) Validate() error {
	if t.Session == nil {
		return errors.New("hace falta tener una session abierta.")
	}
	if t.CreatedAt.IsZero() {
		return errors.New("Fecha vacia o invalida.")
	}

	if len(t.Items) <= 0 {
		return errors.New("El ticket no puede estar vacio.")
	}
	return nil
}

func (t *Ticket) ValidateForCreation() error {
	if t == nil {
		return errors.New("datos de ticket inválidos.")
	}
	if t.Session == nil {
		return errors.New("hace falta tener una session abierta.")
	}
	if t.CreatedAt.IsZero() {
		return errors.New("Fecha vacia o invalida.")
	}
	return nil
}

func (t *Ticket) ValidateForUpdate() error {
	if t.Id <= 0 {
		return errors.New("Ingrese un ID valido")
	}
	if err := t.Validate(); err != nil {
		return err
	}
	return nil
}
