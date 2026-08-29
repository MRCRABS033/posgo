package domain

type Permission struct {
	UserID                       int
	CreateProduct                bool
	DeleteProduct                bool
	ModifiedUnitCostPriceProduct bool
	ModifiedUnitSellPriceProduct bool
	ModifiedDiscount             bool
	ModifiedName                 bool
	ModifiedAvailableDiscount    bool
	ModifiedStock                bool
	ModifiedDepartment           bool

	ModifiedUserName        bool
	ModifieldUserLastName   bool
	ModifiedUserPhoneNumber bool
	ModifiedUserPermissions bool
	ModifiedUserPassword    bool
}
