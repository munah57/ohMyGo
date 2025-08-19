package repository

import (
	"time"
	"vaqua/db"
	"vaqua/models"


)

type TransactionRepository interface {
	GetIncomeByPeriod (UserID uint, start, end time.Time) ([]models.Transaction, error)
	
}

type TransactionRepo struct{}



func (r *TransactionRepo) GetIncomeByPeriod(UserID uint, start, end time.Time) ([]models.Transaction, error) {
    var incomes []models.Transaction
    err := db.Db.Where("user_id = ? AND type = ? AND created_at BETWEEN ? AND ?", UserID, "income", start, end).
        Find(&incomes).Error

    return incomes, err
}

