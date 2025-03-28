package main

import (
	"laliga/database"
	"laliga/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
