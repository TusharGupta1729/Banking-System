package repository

import (
	"banking-system/config"
	"banking-system/models"
)

type BankRepository struct {
}

func NewBankRepository() *BankRepository {
	return &BankRepository{}
}

func (r *BankRepository) Create(bank *models.Bank) error {
	return config.DB.Create(bank).Error
}

func (r *BankRepository) GetAll() ([]models.Bank, error) {
	var banks []models.Bank

	err := config.DB.
		Preload("Branches").
		Find(&banks).Error

	return banks, err
}
