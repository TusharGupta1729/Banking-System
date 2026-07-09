package services

import (
	"errors"

	"banking-system/models"
	"banking-system/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {

	if transaction.AccountID == 0 {
		return errors.New("account id is required")
	}

	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	return s.repo.Create(transaction)
}

func (s *TransactionService) GetTransactions() ([]models.Transaction, error) {
	return s.repo.GetAll()
}
