package handler

import (
	"net/http"
	"time"
	"vaqua/services"

	"github.com/gin-gonic/gin"
) 

type TransactionHandler struct {
	Service *services.TransactionService
}

func (h *TransactionHandler) GetUserIncome(c *gin.Context) {
    userID := c.GetUint("userID") // from auth middleware

    // Placeholder times (no date filtering)
    var start, end time.Time

    incomes, _, err := h.Service.GetIncome(userID, start, end)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch incomes"})
        return
    }

    if len(incomes) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "no transactions found for this user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"incomes": incomes})
}
