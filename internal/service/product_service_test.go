package service

import (
	"errors"
	"testing"

	"product-catalog-api/internal/apperror"
	"product-catalog-api/internal/model"
	"product-catalog-api/internal/repository"
)

type mockProductRepository struct {
	createFunc  func(product *model.Product) error
	getByIDFunc func(id uint) (*model.Product, error)
	updateFunc  func(product *model.Product) error
	deleteFunc  func(id uint) error
	getAllFunc  func(filter repository.ProductFilter) ([]model.Product, error)
}

func (m *mockProductRepository) Create(product *model.Product) error {
	if m.createFunc != nil {
		return m.createFunc(product)
	}

	return nil
}

func (m *mockProductRepository) GetByID(id uint) (*model.Product, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(id)
	}

	return nil, errors.New("not implemented")
}

func (m *mockProductRepository) Update(product *model.Product) error {
	if m.updateFunc != nil {
		return m.updateFunc(product)
	}

	return nil
}

func (m *mockProductRepository) Delete(id uint) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(id)
	}

	return nil
}

func (m *mockProductRepository) GetAll(filter repository.ProductFilter) ([]model.Product, error) {
	if m.getAllFunc != nil {
		return m.getAllFunc(filter)
	}

	return []model.Product{}, nil
}

func TestCreateProduct_Success(t *testing.T) {
	repositoryMock := &mockProductRepository{
		createFunc: func(product *model.Product) error {
			product.ID = 1
			return nil
		},
	}

	productService := NewProductService(repositoryMock)

	product := &model.Product{
		Name:        "iPhone",
		Description: "Smartphone",
		Price:       999.99,
		Category:    "electronics",
	}

	err := productService.CreateProduct(product)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if product.ID != 1 {
		t.Errorf("expected product ID to be 1, got %d", product.ID)
	}
}

func TestCreateProduct_EmptyName(t *testing.T) {
	repositoryMock := &mockProductRepository{}

	productService := NewProductService(repositoryMock)

	product := &model.Product{
		Name:        "",
		Description: "Smartphone",
		Price:       999.99,
		Category:    "electronics",
	}

	err := productService.CreateProduct(product)

	if !errors.Is(err, apperror.ErrProductNameRequired) {
		t.Errorf(
			"expected error %v, got %v",
			apperror.ErrProductNameRequired,
			err,
		)
	}
}

func TestCreateProduct_InvalidPrice(t *testing.T) {
	repositoryMock := &mockProductRepository{}

	productService := NewProductService(repositoryMock)

	product := &model.Product{
		Name:        "iPhone",
		Description: "Smartphone",
		Price:       0,
		Category:    "electronics",
	}

	err := productService.CreateProduct(product)

	if !errors.Is(err, apperror.ErrProductPriceInvalid) {
		t.Errorf(
			"expected error %v, got %v",
			apperror.ErrProductPriceInvalid,
			err,
		)
	}
}

func TestCreateProduct_EmptyCategory(t *testing.T) {
	repositoryMock := &mockProductRepository{}

	productService := NewProductService(repositoryMock)

	product := &model.Product{
		Name:        "iPhone",
		Description: "Smartphone",
		Price:       999.99,
		Category:    "",
	}

	err := productService.CreateProduct(product)

	if !errors.Is(err, apperror.ErrProductCategoryEmpty) {
		t.Errorf(
			"expected error %v, got %v",
			apperror.ErrProductCategoryEmpty,
			err,
		)
	}
}

func TestGetProductByID_NotFound(t *testing.T) {
	repositoryMock := &mockProductRepository{
		getByIDFunc: func(id uint) (*model.Product, error) {
			return nil, errors.New("product not found")
		},
	}

	productService := NewProductService(repositoryMock)

	product, err := productService.GetProductByID(999)

	if product != nil {
		t.Errorf("expected product to be nil")
	}

	if !errors.Is(err, apperror.ErrProductNotFound) {
		t.Errorf(
			"expected error %v, got %v",
			apperror.ErrProductNotFound,
			err,
		)
	}
}
