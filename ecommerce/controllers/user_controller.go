package controllers

import (
	"ecommerce/database"
	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	// Hide passwords in the response
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(200, users)
}

// Get a single user
func GetUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Hide the password in the response
	user.Password = ""

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Save the user
	database.DB.Create(&user)
	c.JSON(201, user)
}

// Update a user
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&user)
	c.JSON(200, user)
}

// Delete a user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Delete the user and cascade delete their products
	database.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User and their products deleted successfully"})
}
