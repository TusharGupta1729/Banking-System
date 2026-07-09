package repository

import (
	"banking-system/config"
	"banking-system/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return config.DB.Create(transaction).Error
}

func (r *TransactionRepository) GetAll() ([]models.Transaction, error) {

	var transactions []models.Transaction

	err := config.DB.
		Preload("Account").
		Find(&transactions).Error

	return transactions, err
}

func (r *TransactionRepository) CreateWithTransaction(
	tx *gorm.DB,
	transaction *models.Transaction,
) error {
	return tx.Create(transaction).Error
}
