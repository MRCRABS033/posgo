package domain

import (
	"errors"
	"strings"
)

type CashOut struct {
	ID        int
	userID    int
	sessionID int
	Quantity  float64
	Concept   string
}

func (c *CashOut) Validate() error {
	if c == nil {
		return errors.New("datos de la salida de efectivo invalido.")
	}
	if c.Quantity <= 0 {
		return errors.New("Cantidad invalida, por favor ingrese un monto mayor a 0.00$")
	}
	if strings.TrimSpace(c.Concept) == "" {
		return errors.New("Por favor escriba el motivo de la salida de efectivo.")
	}
	return nil
}
