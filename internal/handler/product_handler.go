package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"product-catalog-api/internal/apperror"
	"product-catalog-api/internal/model"
	"product-catalog-api/internal/repository"
	"product-catalog-api/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

type createProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

type updateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request createProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	product := model.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Category:    request.Category,
	}

	err := h.service.CreateProduct(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	product, err := h.service.GetProductByID(uint(id))

	if err != nil {
		if errors.Is(err, apperror.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	var request updateProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	product := model.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Category:    request.Category,
	}

	updatedProduct, err := h.service.UpdateProduct(uint(id), &product)

	if err != nil {
		if errors.Is(err, apperror.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	err = h.service.DeleteProduct(uint(id))

	if err != nil {
		if errors.Is(err, apperror.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	filter := repository.ProductFilter{
		Category: c.Query("category"),
	}

	if limitString := c.Query("limit"); limitString != "" {
		limit, err := strconv.Atoi(limitString)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid limit",
			})
			return
		}

		filter.Limit = limit
	}

	if offsetString := c.Query("offset"); offsetString != "" {
		offset, err := strconv.Atoi(offsetString)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid offset",
			})
			return
		}

		filter.Offset = offset
	}

	if priceFromString := c.Query("price_from"); priceFromString != "" {
		priceFrom, err := strconv.ParseFloat(priceFromString, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid price_from",
			})
			return
		}

		if priceFrom < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "price_from cannot be negative",
			})
			return
		}

		filter.PriceFrom = &priceFrom
	}

	if priceToString := c.Query("price_to"); priceToString != "" {
		priceTo, err := strconv.ParseFloat(priceToString, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid price_to",
			})
			return
		}

		if priceTo < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "price_to cannot be negative",
			})
			return
		}

		filter.PriceTo = &priceTo
	}

	if filter.PriceFrom != nil && filter.PriceTo != nil {
		if *filter.PriceFrom > *filter.PriceTo {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "price_from cannot be greater than price_to",
			})
			return
		}
	}

	products, err := h.service.GetProducts(filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
