package controllers

import (
	"net/http"
	"strconv"

	"laliga/database"
	"laliga/models"

	"github.com/gin-gonic/gin"
)

// Obtener todos los partidos
func GetAllMatches(c *gin.Context) {
	var matches []models.Match
	database.DB.Find(&matches)
	c.JSON(http.StatusOK, matches)
}

// Obtener un partido por ID
func GetMatchByID(c *gin.Context) {
	id := c.Param("id")
	var match models.Match
	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	c.JSON(http.StatusOK, match)
}

// Crear un nuevo partido
func CreateMatch(c *gin.Context) {
	var match models.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&match)
	c.JSON(http.StatusCreated, match)
}

// Actualizar un partido existente
func UpdateMatch(c *gin.Context) {
	id := c.Param("id")
	var match models.Match
	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	var input models.Match
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match.HomeTeam = input.HomeTeam
	match.AwayTeam = input.AwayTeam
	match.MatchDate = input.MatchDate

	database.DB.Save(&match)
	c.JSON(http.StatusOK, match)
}

// Eliminar un partido
func DeleteMatch(c *gin.Context) {
	id := c.Param("id")
	idNum, _ := strconv.Atoi(id)
	result := database.DB.Delete(&models.Match{}, idNum)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	c.Status(http.StatusNoContent)
}
