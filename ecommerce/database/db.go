package database

import (
	"ecommerce/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase() {
	// Replace with your MySQL DSN
	dsn := "ecommerce_user:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

// Migrate applies the schema migrations for the database
func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Product{})
}
