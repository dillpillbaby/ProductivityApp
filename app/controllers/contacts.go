package controllers

import (
	"net/http"

	"productivityapp/models"

	"github.com/gin-gonic/gin"
)

// This comes from the request to verify that we have the right stuff
type CreateContactInput struct {
	Name         string `json:"name" binding:"required"`
	Phone_Number string `json:"phoneNumber" binding:"required"`
}

func FindContacts(c *gin.Context) {
	var contacts []models.Contact
	models.DB.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

func CreateContact(c *gin.Context) {
	var input CreateContactInput
	// Verify we have the right fields and throw error if not
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Creating Contact with the input 
	contact := models.Contact{Name: input.Name, Phone_Number: input.Phone_Number}
	models.DB.Create(&contact)

	// Send back what we have
	c.JSON(http.StatusOK, gin.H{"data": contact})
}
