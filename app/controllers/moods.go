package controllers

import (
	"net/http"

	"productivityapp/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

// This comes from the request to verify that we have the right stuff
type CreateMoodInput struct {
	Name            string `json:"name" binding:"required"`
	PositivityLevel int    `json:"positivityLevel" binding:"required"`
	ColorID         int    `json:"colorId" binding:"required"`
}

func FindMoods(c *gin.Context) {
	var moods []models.Mood
	if err := models.DB.Preload("Color").Find(&moods).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": moods})
}

func DeleteMood(c *gin.Context) {
	var mood models.Mood
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	models.DB.Where("id = ?", id).Delete(&mood)
	models.DB.Delete(&mood)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func CreateMood(c *gin.Context) {
	var input CreateMoodInput
	// Verify we have the right fields and throw error if not
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Creating Contact with the input
	var color models.Color
	if err := models.DB.First(&color, input.ColorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Color not found"})
		return
	}

	// Create the mood
	mood := models.Mood{
		Name:            input.Name,
		ColorID:         int(color.ID),
		Color:           color,
		PositivityLevel: input.PositivityLevel,
	}

	if err := models.DB.Create(&mood).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send back what we have
	c.JSON(http.StatusOK, gin.H{"data": mood})
}
