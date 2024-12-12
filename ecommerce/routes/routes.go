package routes

import (
    "ecommerce/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    userRoutes := router.Group("/users")
    {
        userRoutes.GET("/", controllers.GetUsers)
        userRoutes.POST("/", controllers.CreateUser)
    }

    productRoutes := router.Group("/products")
    {
        productRoutes.GET("/", controllers.GetProducts)
        productRoutes.POST("/", controllers.CreateProduct)
    }
}
