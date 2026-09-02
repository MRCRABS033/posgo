package domain

import (
	"errors"
	"strings"
)

type Department struct {
	ID   int
	Name string
}

func (d *Department) Validate() error {
	if strings.TrimSpace(d.Name) == "" {
		return errors.New("El nombre del departamento no puede venir vacio.")
	}
	return nil
}

func (d *Department) ValidateForUpdate() error {
	if d.ID <= 0 {
		return errors.New("El ID no puede estar vacio.")
	}
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}
