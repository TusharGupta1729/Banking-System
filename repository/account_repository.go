package repository

import (
	"banking-system/config"
	"banking-system/models"
)

type AccountRepository struct {
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) Create(account *models.Account) error {
	return config.DB.Create(account).Error
}

func (r *AccountRepository) GetAll() ([]models.Account, error) {

	var accounts []models.Account

	err := config.DB.Find(&accounts).Error

	return accounts, err
}
