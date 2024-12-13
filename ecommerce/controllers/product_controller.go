package controllers

import (
	"ecommerce/database"
	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(200, products)
}

// Get a single product
func GetProduct(c *gin.Context) {
	var product models.Product
	if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind JSON ke product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Periksa apakah User dengan ID yang diberikan ada
	var user models.User
	if err := database.DB.First(&user, product.UserID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Simpan produk ke database
	database.DB.Create(&product)
	c.JSON(201, product)
}

// Update a product
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&product)
	c.JSON(200, product)
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	database.DB.Delete(&product)
	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}

func GetProductsByUser(c *gin.Context) {
	var products []models.Product
	userID := c.Param("user_id")

	if err := database.DB.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		c.JSON(404, gin.H{"error": "Products not found"})
		return
	}

	c.JSON(200, products)
}
