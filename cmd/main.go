package main

import (
	"github.com/gin-gonic/gin"

	"product-catalog-api/internal/database"
)

func main() {
	db := database.Connect()

	_ = db

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run(":8080")
}