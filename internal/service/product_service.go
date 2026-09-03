package service

import (
	"strings"

	"product-catalog-api/internal/apperror"
	"product-catalog-api/internal/model"
	"product-catalog-api/internal/repository"
)

type ProductService struct {
	repository repository.ProductRepositoryInterface
}

func NewProductService(repository repository.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	product.Name = strings.TrimSpace(product.Name)
	product.Description = strings.TrimSpace(product.Description)
	product.Category = strings.TrimSpace(product.Category)

	if product.Name == "" {
		return apperror.ErrProductNameRequired
	}

	if product.Price <= 0 {
		return apperror.ErrProductPriceInvalid
	}

	if product.Category == "" {
		return apperror.ErrProductCategoryEmpty
	}

	return s.repository.Create(product)
}

func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	product, err := s.repository.GetByID(id)

	if err != nil {
		return nil, apperror.ErrProductNotFound
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(id uint, product *model.Product) (*model.Product, error) {
	existingProduct, err := s.repository.GetByID(id)

	if err != nil {
		return nil, apperror.ErrProductNotFound
	}

	product.Name = strings.TrimSpace(product.Name)
	product.Description = strings.TrimSpace(product.Description)
	product.Category = strings.TrimSpace(product.Category)

	if product.Name == "" {
		return nil, apperror.ErrProductNameRequired
	}

	if product.Price <= 0 {
		return nil, apperror.ErrProductPriceInvalid
	}

	if product.Category == "" {
		return nil, apperror.ErrProductCategoryEmpty
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
		return apperror.ErrProductNotFound
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
