package services

import "vaqua/repository"

type TransactionService struct {
	Repo repository.TransactionRepository
}