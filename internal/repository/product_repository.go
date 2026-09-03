package repository

import (
	"errors"

	"gorm.io/gorm"

	"product-catalog-api/internal/model"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product

	err := r.db.First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

type ProductFilter struct {
	Limit     int
	Offset    int
	Category  string
	PriceFrom *float64
	PriceTo   *float64
}

func (r *ProductRepository) GetAll(filter ProductFilter) ([]model.Product, error) {
	var products []model.Product

	query := r.db.Model(&model.Product{})

	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}

	if filter.PriceFrom != nil {
		query = query.Where("price >= ?", *filter.PriceFrom)
	}

	if filter.PriceTo != nil {
		query = query.Where("price <= ?", *filter.PriceTo)
	}

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	err := query.Order("id DESC").Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}