package usecase

import (
	"errors"
	"posgo/internal/domain"
)

type ProductUseCase struct {
	repo domain.ProductRepository
}

func NewProductUseCase(repo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

// get a product
func (uc *ProductUseCase) GetByCode(code string) (*domain.Product, error) {
	if code == "" {
		return nil, errors.New("Por favor ingrese un codigo.")
	}

	product, err := uc.repo.GetProductByCode(code)

	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("Producto no encontrado")
	}

	return product, nil
}

func (uc *ProductUseCase) GetProductLowStock() ([]*domain.Product, error) {
	products, err := uc.repo.GetProductLowStock()
	if err != nil {
		return nil, err
	}
	if products != nil {
		return nil, errors.New("Sin productos con bajo stock")
	}

	return products, nil
}

func (uc *ProductUseCase) GetProducByDepartment(department string) ([]*domain.Product, error) {
	products, err := uc.repo.GetProductbyDepartment(department)

	if err != nil {
		return nil, err
	}
	if products != nil {
		return nil, errors.New("No se econtraron productos en este departamento.")
	}

	return products, nil
}

// create a new product
func (uc *ProductUseCase) CreateProduct(code, name string, costPrice, sellPrice float64, stock int, department *domain.Department) (*domain.Product, error) {

	if sellPrice < costPrice {
		return nil, errors.New("el precio de venta no puede ser menor que el costo.")
	}

	productExist, err := uc.repo.GetProductByCode(code)

	if productExist != nil {
		return nil, errors.New("ya existe un producto registrado con ese codigo.")
	}

	if department == nil {
		department = &domain.Department{
			ID:   0,
			Name: "Sin departamento",
		}
	}

	product := &domain.Product{
		Code:              code,
		Name:              name,
		UnitCostPrice:     costPrice,
		UnitSellPrice:     sellPrice,
		Discount:          0.00,
		IsSingleProduct:   true,
		AvailableDiscount: false,
		Stock:             stock,
		Available:         true,
		Department:        *department,
	}

	err = uc.repo.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUseCase) UpdateProduct(code, name string, unitCost, unitSell, discount float64, isSingleProduct, availableDiscount bool, stock int, available bool, department *domain.Department) error {

	existingProduct, err := uc.repo.GetProductByCode(code)
	if err != nil {
		return errors.New("El producto que intentas modificar no existe.")
	}

	if unitSell < unitCost {
		return errors.New("el precio de venta no puede ser menor que el costo.")
	}

	existingProduct.Name = name
	existingProduct.UnitCostPrice = unitCost
	existingProduct.UnitSellPrice = unitSell
	existingProduct.Discount = discount
	existingProduct.IsSingleProduct = isSingleProduct
	existingProduct.AvailableDiscount = availableDiscount
	existingProduct.Available = available
	existingProduct.Stock = stock
	existingProduct.Department = *department

	return uc.repo.UpdateProduct(existingProduct)
}
