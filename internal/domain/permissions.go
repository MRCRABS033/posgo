package domain

import "errors"

type Permission struct {
	ID                           int
	UserID                       int
	CreateProduct                bool
	DeleteProduct                bool
	ModifiedUnitCostPriceProduct bool
	ModifiedUnitSellPriceProduct bool
	ModifiedStockProduct         bool
	ModifiedDiscount             bool
	ModifiedName                 bool
	ModifiedAvailableDiscount    bool
	ModifiedDepartment           bool
	ModifiedUserName             bool
	ModifiedUserLastName         bool
	ModifiedUserPhoneNumber      bool
	ModifiedUserPermissions      bool
	ModifiedUserPassword         bool
}

func (p *Permission) Validate() error {
	if p == nil {
		return errors.New("Datos de permisos invalidos.")
	}

	if p.UserID <= 0 {
		return errors.New("El ID del usuario asociado a los permisos no es valido.")
	}

	return nil
}
