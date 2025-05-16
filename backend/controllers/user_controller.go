package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Bind the request body to the LoginDTO struct
	var loginDTO domain.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the Login function from the services package
	token, err := services.Login(loginDTO.Username, loginDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Return the token in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}
