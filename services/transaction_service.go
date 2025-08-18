package services

import (
	"time"
	"vaqua/models"
	"vaqua/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
}

func (s *TransactionService) CreateTransaction(tx *models.Transaction) error {
	return s.Repo.CreateTransaction(tx)
}

func (s *TransactionService) GetExpensesByUser(userID uint, fromDate time.Time) ([]models.Transaction, error) {
	return s.Repo.GetExpensesByUser(userID, fromDate)
}

func (s *TransactionService) GetExpenseSummaryByUser(userID uint, fromDate time.Time) ([]repository.ExpenseSummary, error) {
	return s.Repo.GetExpenseSummaryByUser(userID, fromDate)
}
