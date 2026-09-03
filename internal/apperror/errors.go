package apperror

import "errors"

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductNameRequired  = errors.New("product name is required")
	ErrProductPriceInvalid  = errors.New("product price must be greater than zero")
	ErrProductCategoryEmpty = errors.New("product category is required")
)
