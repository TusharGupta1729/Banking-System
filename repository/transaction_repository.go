package repository

import (
	"banking-system/config"
	"banking-system/models"
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

	err := config.DB.Find(&transactions).Error

	return transactions, err
}
