package domain

import (
	"errors"
	"strings"
)

type CashIn struct {
	ID        int
	UserID    int
	SessionID int
	Concept   string
	Quantity  float64
}

func (c *CashIn) Validate() error {
	if c == nil {
		return errors.New("datos de la entrada de efectivo invalidos.")
	}
	if c.Quantity <= 0 {
		return errors.New("Cantidad Invalida, por favor ingrese un monto mayor a 0.00$")
	}
	if strings.TrimSpace(c.Concept) == "" {
		return errors.New("Por favor escriba el motivo de entrada de efectivo.")
	}
	return nil
}
