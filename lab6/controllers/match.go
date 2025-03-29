package controllers

import (
	"net/http"
	"strconv"

	"lab6/database"
	"lab6/models"

	"github.com/gin-gonic/gin"
)

// GetMatches obtiene todos los partidos registrados
// @Router /matches [get]
func GetMatches(c *gin.Context) {
	var matches []models.Match
	database.DB.Find(&matches)
	c.JSON(http.StatusOK, matches)
}

// GetMatchByID obtiene un partido por ID
// @Router /matches/{id} [get]
func GetMatchByID(c *gin.Context) {
	id := c.Param("id")
	var match models.Match
	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match was not found"})
		return
	}
	c.JSON(http.StatusOK, match)
}

// CreateMatch crea un nuevo partido
// @Router /matches [post]
func CreateMatch(c *gin.Context) {
	var match models.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&match)
	c.JSON(http.StatusCreated, match)
}

// UpdateMatch actualiza un partido existente
// @Router /matches/{id} [put]
func UpdateMatch(c *gin.Context) {
	id := c.Param("id")
	var match models.Match
	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match was not found"})
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

// DeleteMatch elimina un partido
// @Router /matches/{id} [delete]
func DeleteMatch(c *gin.Context) {
	id := c.Param("id")
	idNum, _ := strconv.Atoi(id)
	result := database.DB.Delete(&models.Match{}, idNum)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match was not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateGoals registra un gol en el partido
// @Router /matches/{id}/goals [patch]
func UpdateGoals(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match was not found"})
		return
	}

	match.TotalGoals++

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error at trying to save the match"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// AddYellowCard registra una tarjeta amarilla
// @Router /matches/{id}/yellowcards [patch]
func AddYellowCard(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match was not found"})
		return
	}

	match.YellowCards++

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error at trying to save the match"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// AddRedCard registra una tarjeta roja
// @Router /matches/{id}/redcards [patch]
func AddRedCard(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match was not found"})
		return
	}

	match.RedCards++

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error at trying to save the match"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// AddExtraTime registra tiempo extra en el partido
// @Router /matches/{id}/extratime [patch]
func AddExtraTime(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match was not found"})
		return
	}

	match.ExtraTime = true

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error at trying to save the match"})
		return
	}

	c.JSON(http.StatusOK, match)
}
