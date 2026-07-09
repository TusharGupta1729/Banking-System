package repository

import (
	"banking-system/config"
	"banking-system/models"
	"gorm.io/gorm"
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

	err := config.DB.
		Preload("Customer").
		Preload("Branch").
		Preload("Transactions").
		Find(&accounts).Error

	return accounts, err
}

func (r *AccountRepository) GetByID(id uint) (*models.Account, error) {

	var account models.Account

	err := config.DB.
		Preload("Customer").
		Preload("Branch").
		Preload("Transactions").
		First(&account, id).Error

	return &account, err
}

func (r *AccountRepository) Update(account *models.Account) error {
	return config.DB.Save(account).Error
}

func (r *AccountRepository) UpdateWithTransaction(
	tx *gorm.DB,
	account *models.Account,
) error {

	return tx.Save(account).Error
}
