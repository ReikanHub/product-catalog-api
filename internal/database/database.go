package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"product-catalog-api/internal/model"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&model.Product{})

	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("database connected successfully")

	return db
}
