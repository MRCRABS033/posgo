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

func (uc *ProductUseCase) GetProductByCode(code string) (*domain.Product, error) {
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

func (uc *ProductUseCase) GetProductByName(name string) (*domain.Product, error) {
	if name == "" {
		return nil, errors.New("Por favor ingrese un nombre.")
	}

	product, err := uc.repo.GetProductByName(name)

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
	if len(products) == 0 {
		return nil, errors.New("Sin productos con bajo stock")
	}

	return products, nil
}

func (uc *ProductUseCase) GetProducByDepartment(department string) ([]*domain.Product, error) {
	products, err := uc.repo.GetProductbyDepartment(department)

	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("No se econtraron productos en este departamento.")
	}

	return products, nil
}

//---------------------------------------------------------------

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
		return err
	}

	if existingProduct == nil {
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

func (uc *ProductUseCase) RemoveProduct(code string) error {
	existProduct, err := uc.repo.GetProductByCode(code)
	if err != nil {
		return err
	}
	if existProduct == nil {
		return errors.New("el producto que intentas borrar no existe.")
	}

	return uc.repo.RemoveProduct(existProduct.Code)
}
