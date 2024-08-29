package controllers

import (
	"net/http"

	"productivityapp/models"

	"github.com/gin-gonic/gin"
)

// This comes from the request to verify that we have the right stuff
type CreateColorInput struct {
	Name         string `json:"name" binding:"required"`
	ColorCode string `json:"colorCode" binding:"required"`
}

func FindColors(c *gin.Context) {
	var colors []models.Color
	models.DB.Find(&colors)

	c.JSON(http.StatusOK, gin.H{"data": colors})
}

func CreateColor(c *gin.Context) {
	var input CreateColorInput
	// Verify we have the right fields and throw error if not
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Creating Contact with the input
	color := models.Color{Name: input.Name, ColorCode: input.ColorCode}
	models.DB.Create(&color)

	// Send back what we have
	c.JSON(http.StatusOK, gin.H{"data": color})
}
