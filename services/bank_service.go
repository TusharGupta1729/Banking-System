package services

import (
	"errors"

	"banking-system/models"
	"banking-system/repository"
)

type BankService struct {
	repo *repository.BankRepository
}

func NewBankService(repo *repository.BankRepository) *BankService {
	return &BankService{
		repo: repo,
	}
}

func (s *BankService) CreateBank(bank *models.Bank) error {

	if bank.Name == "" {
		return errors.New("bank name is required")
	}

	return s.repo.Create(bank)
}

func (s *BankService) GetBanks() ([]models.Bank, error) {
	return s.repo.GetAll()
}
