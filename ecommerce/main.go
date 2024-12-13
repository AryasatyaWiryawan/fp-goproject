package main

import (
    "ecommerce/database"
    "ecommerce/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database connection
    database.InitDatabase()

    // Migrate the models
    database.Migrate()

    // Set up Gin routes
    router := gin.Default()
	router.Static("/public", "./public")
    routes.SetupRoutes(router)

    // Start the application
    router.Run(":8080")
}
