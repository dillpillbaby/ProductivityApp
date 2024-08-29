package controllers

import (
	"fmt"
	"net/http"

	"productivityapp/models"

	"github.com/gin-gonic/gin"
)

// This comes from the request to verify that we have the right stuff
type CreateJournalEntryInput struct {
	Description string `json:"description" binding:"required"`
	MoodID int `json:"moodId" binding:"required"`
}

func FindJournalEntries(c *gin.Context) {
	var journal_entries []models.JournalEntry
	if err := models.DB.Preload("Mood.Color").Find(&journal_entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": journal_entries})
}

func CreateJournalEntry(c *gin.Context) {
	var input CreateJournalEntryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	fmt.Println("Input", input)
	fmt.Println(input.Description)
	fmt.Println(input.MoodID)
	var mood models.Mood
	if err := models.DB.First(&mood, input.MoodID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mood not found"})
		return
	}
	fmt.Println("Mood", mood)

	// Create the journal entry
	journal_entry := models.JournalEntry{
		Description: input.Description,
		MoodID:      int(mood.ID),
		Mood: mood,
	}

	if err := models.DB.Create(&journal_entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": journal_entry})
}
