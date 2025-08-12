package handler

import (
	"net/http"
	"strconv"
	"vaqua/models"
	"vaqua/services"

	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	Service services.TransferService
}

func (h *TransferHandler) CreateTransfer(c *gin.Context) {
	// Bind JSON body to the TransferRequest model
	var request models.TransferRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set in middleware)
	userIDStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// userIDStr is interface{}, assert it to string
	userIDStrVal, ok := userIDStr.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID in context is not a string"})
		return
	}

	// Convert userID string to uint64
	userID, err := strconv.ParseUint(userIDStrVal, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call service to perform transfer
	err = h.Service.TransferFunds(uint(userID), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}