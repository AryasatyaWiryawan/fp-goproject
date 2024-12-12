package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// User Routes
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)    // Update user
		userRoutes.DELETE("/:id", controllers.DeleteUser) // Delete user
	}

	// Product Routes
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.PUT("/:id", controllers.UpdateProduct)    // Update product
		productRoutes.DELETE("/:id", controllers.DeleteProduct) // Delete product
	}
}
