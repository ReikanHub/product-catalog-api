package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"product-catalog-api/internal/database"
	"product-catalog-api/internal/handler"
	"product-catalog-api/internal/middleware"
	"product-catalog-api/internal/repository"
	"product-catalog-api/internal/service"
)

func main() {
	db := database.Connect()

	productRepository := repository.NewProductRepository(db)

	productService := service.NewProductService(productRepository)

	productHandler := handler.NewProductHandler(productService)

	router := gin.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	log.Println("server started on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
