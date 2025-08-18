package repository

import (
	"time"
	"vaqua/db"
	"vaqua/models"
)

type ExpenseSummary struct {
	Description string  `json:"description"`
	Total       float64 `json:"total"`
}

type TransactionRepository interface {
	CreateTransaction(tx *models.Transaction) error
	GetExpensesByUser(userID uint, fromDate time.Time) ([]models.Transaction, error)
	GetExpenseSummaryByUser(userID uint, fromDate time.Time) ([]ExpenseSummary, error)
}

type TransactionRepo struct{}

func (r *TransactionRepo) CreateTransaction(tx *models.Transaction) error {
	return db.Db.Create(tx).Error
}

func (r *TransactionRepo) GetExpensesByUser(userID uint, fromDate time.Time) ([]models.Transaction, error) {
	var expenses []models.Transaction
	q := db.Db.Where("user_id = ? AND type = ?", userID, "expense")
	if !fromDate.IsZero() {
		q = q.Where("created_at >= ?", fromDate)
	}
	if err := q.Order("created_at DESC").Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

// Group and sum expenses by description (works as “category”)
func (r *TransactionRepo) GetExpenseSummaryByUser(userID uint, fromDate time.Time) ([]ExpenseSummary, error) {
	var rows []ExpenseSummary
	q := db.Db.Model(&models.Transaction{}).
		Select("description, SUM(amount) as total").
		Where("user_id = ? AND type = ?", userID, "expense")
	if !fromDate.IsZero() {
		q = q.Where("created_at >= ?", fromDate)
	}
	if err := q.Group("description").Order("total DESC").Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
