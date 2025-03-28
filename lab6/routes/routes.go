package routes

import (
	"laliga/controllers"

	"github.com/gin-gonic/gin"
)

// Configurar las rutas de la API
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/matches", controllers.GetAllMatches)
		api.GET("/matches/:id", controllers.GetMatchByID)
		api.POST("/matches", controllers.CreateMatch)
		api.PUT("/matches/:id", controllers.UpdateMatch)
		api.DELETE("/matches/:id", controllers.DeleteMatch)
	}
}
