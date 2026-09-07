package ui

import "github.com/charmbracelet/bubbles/textinput"

type CreateProduct struct {
	Code              string
	ProductName       string
	UnitCostPrice     float64
	UnitSellPrice     float64
	Discount          float64
	IsSingleProduct   bool
	AvailableDiscount bool
	Stock             float64
	Available         bool
	DepartmentName    string
}

type ProductModel struct {
	codeInput          textinput.Model
	productNameInput   textinput.Model
	unitCostPriceInput textinput.Model
}
