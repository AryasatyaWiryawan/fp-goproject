package main

import (
	"ecommerce/database"
	"ecommerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	database.Migrate()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
