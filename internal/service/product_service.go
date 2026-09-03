package service

import (
	"errors"
	"strings"

	"product-catalog-api/internal/model"
	"product-catalog-api/internal/repository"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductService struct {
	repository *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	product.Name = strings.TrimSpace(product.Name)
	product.Description = strings.TrimSpace(product.Description)
	product.Category = strings.TrimSpace(product.Category)

	if product.Name == "" {
		return errors.New("product name is required")
	}

	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}

	if product.Category == "" {
		return errors.New("product category is required")
	}

	return s.repository.Create(product)
}

func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	product, err := s.repository.GetByID(id)

	if err != nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(id uint, product *model.Product) (*model.Product, error) {
	existingProduct, err := s.repository.GetByID(id)

	if err != nil {
		return nil, ErrProductNotFound
	}

	product.Name = strings.TrimSpace(product.Name)
	product.Description = strings.TrimSpace(product.Description)
	product.Category = strings.TrimSpace(product.Category)

	if product.Name == "" {
		return nil, errors.New("product name is required")
	}

	if product.Price <= 0 {
		return nil, errors.New("product price must be greater than zero")
	}

	if product.Category == "" {
		return nil, errors.New("product category is required")
	}

	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.Category = product.Category

	err = s.repository.Update(existingProduct)

	if err != nil {
		return nil, err
	}

	return existingProduct, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.repository.GetByID(id)

	if err != nil {
		return ErrProductNotFound
	}

	err = s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) GetProducts(filter repository.ProductFilter) ([]model.Product, error) {
	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	if filter.Offset < 0 {
		filter.Offset = 0
	}

	return s.repository.GetAll(filter)
}