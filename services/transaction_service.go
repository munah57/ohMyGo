package services

import (
	"time"
	"vaqua/models"

	"vaqua/repository"
	

)
type TransactionServices interface {
	GetIncome (userID uint, start, end time.Time) ([]models.Transaction, float64, error) 
}

type TransactionService struct {
	Repo *repository.TransactionRepo
}


func NewTransactionService(repo *repository.TransactionRepo) *TransactionService {
	return &TransactionService{Repo: repo}
}

func (s *TransactionService) GetIncome(userID uint, start, end time.Time) ([]models.Transaction, float64, error) {
	incomes, err := s.Repo.GetIncomeByPeriod(userID, start, end)
	if err != nil {
		return nil, 0, err
	}

	var total float64
	for _, t := range incomes {
		total += t.Amount
	}
	return incomes, total, nil
}
