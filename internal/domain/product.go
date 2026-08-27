package domain

type Product struct {
	ID                int
	Code              string
	Name              string
	UnitCostPrice     float64
	UnitSellPrice     float64
	Discount          float64
	IsSingleProduct   bool
	AvailableDiscount bool
	Stock             int
	Available         bool
	Department        Department
}
