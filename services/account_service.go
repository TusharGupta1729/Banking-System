package services

import (
	"errors"

	"banking-system/models"
	"banking-system/repository"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) CreateAccount(account *models.Account) error {

	if account.AccountNumber == "" {
		return errors.New("account number is required")
	}

	if account.CustomerID == 0 {
		return errors.New("customer id is required")
	}

	if account.BranchID == 0 {
		return errors.New("branch id is required")
	}

	return s.repo.Create(account)
}

func (s *AccountService) GetAccounts() ([]models.Account, error) {
	return s.repo.GetAll()
}
