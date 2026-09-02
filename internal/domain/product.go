package domain

import (
	"errors"
	"strings"
)

type Product struct {
	ID                int
	Code              string
	Name              string
	UnitCostPrice     float64
	UnitSellPrice     float64
	Discount          float64
	IsSingleProduct   bool
	AvailableDiscount bool
	Stock             float64
	Available         bool
	Department        *Department
}

func (p *Product) Validate() error {

	if p == nil {
		return errors.New("datos del producto inválidos.")
	}

	if strings.TrimSpace(p.Name) == "" {
		return errors.New("Ingrese un nombre valido.")
	}

	if p.UnitCostPrice <= 0 {
		return errors.New("El precio de costo no puede ser menor o igual a 0.00$")
	}

	if p.UnitSellPrice <= 0 {
		return errors.New("El precio de venta no puede ser menor o igual a 0.00$")
	}

	if p.UnitSellPrice <= p.UnitCostPrice {
		return errors.New("El precio de venta no puede ser menor a precio de costo.")
	}

	if p.Stock <= 0 {
		return errors.New("El stock no puede ser menor o igual a 0.00")
	}

	return nil
}

func (p *Product) ValidateForUpdate() error {
	if p.ID <= 0 {
		return errors.New("El ID no puede estar vacio")
	}
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}
