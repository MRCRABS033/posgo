package domain

import (
	"errors"
)

type TicketItem struct {
	Session  *Session
	Product  *Product
	Quantity float64
	Subtotal float64
}

func (t *TicketItem) Validate() error {
	if t.Session == nil {
		return errors.New("hace falta tener una sesion abierta.")
	}

	if t.Product == nil {
		return errors.New("El producto no puede estar vacio.")
	}

	if t.Quantity <= 0.0000 {
		return errors.New("La cantidad no puede ser menor a 0.0000")
	}

	if t.Subtotal <= 0.00 {
		return errors.New("El total no puede ser menor a 0.00.")
	}
	return nil
}

func (t *TicketItem) AddQuantity(amount float64) error {
	if amount <= 0 {
		return errors.New("La cantidad debe ser mayor a cero.")
	}
	if (t.Quantity + amount) > t.Product.Stock {
		return errors.New("La cantidad no puede ser mayor al stock que tiene el producto.")
	}

	t.Quantity += amount

	if t.Product != nil {
		t.Subtotal = t.Quantity * t.Product.UnitSellPrice
	}
	return nil
}

func (t *TicketItem) ReduceQuantity(amount float64) error {
	if amount <= 0 {
		return errors.New("La cantidad debe ser mayor a cero.")
	}

	if amount > t.Quantity {
		return errors.New("No puede reducir una cantidad mayor a la que tiene item en stock")
	}
	t.Quantity -= amount

	if t.Product != nil {
		t.Subtotal = t.Quantity * t.Product.UnitSellPrice
	}

	return t.Validate()
}
