package controllers

import (
    "ecommerce/database"
    "ecommerce/models"
    "github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
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

	database.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}