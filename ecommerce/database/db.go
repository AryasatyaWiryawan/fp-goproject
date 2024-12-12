package database

import (
	"ecommerce/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Product{})
}
