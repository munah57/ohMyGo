package handler

import (
	"net/http"
	"time"
	"vaqua/models"
	"vaqua/services"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	Service *services.TransactionService
}

// POST 
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var tx models.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uidAny, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	tx.UserID = uidAny.(uint)

	if tx.Type != "expense" && tx.Type != "income" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type must be 'expense' or 'income'"})
		return
	}
	if tx.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be > 0"})
		return
	}

	if err := h.Service.CreateTransaction(&tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}
	c.JSON(http.StatusCreated, tx)
}

// GET 
func (h *TransactionHandler) GetExpenses(c *gin.Context) {
	uidAny, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID := uidAny.(uint)

	fromDate, bad := parsePeriod(c.Query("period"))
	if bad {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid period; use 1m, 6m, or 1y"})
		return
	}

	expenses, err := h.Service.GetExpensesByUser(userID, fromDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch expenses"})
		return
	}
	if len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no expenses found"})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

// GET 
func (h *TransactionHandler) GetExpenseSummary(c *gin.Context) {
	uidAny, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID := uidAny.(uint)

	fromDate, bad := parsePeriod(c.Query("period"))
	if bad {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid period; use 1m, 6m, or 1y"})
		return
	}

	rows, err := h.Service.GetExpenseSummaryByUser(userID, fromDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch expense summary"})
		return
	}
	if len(rows) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no expense summary for this period"})
		return
	}
	c.JSON(http.StatusOK, rows)
}

// Helper
func parsePeriod(p string) (time.Time, bool) {
	switch p {
	case "1m":
		return time.Now().AddDate(0, -1, 0), false
	case "6m":
		return time.Now().AddDate(0, -6, 0), false
	case "1y":
		return time.Now().AddDate(-1, 0, 0), false
	case "", "all":
		return time.Time{}, false // zero time = no filter
	default:
		return time.Time{}, true
	}
}
