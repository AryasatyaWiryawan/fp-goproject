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
		userRoutes.GET("/:id", controllers.GetUser) // Add this line
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)    // Update user
		userRoutes.DELETE("/:id", controllers.DeleteUser) // Delete user
	}

	// Product Routes
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.GET("/:id", controllers.GetProduct)
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
		productRoutes.GET("/user/:user_id", controllers.GetProductsByUser) // Produk berdasarkan User
	}

}
