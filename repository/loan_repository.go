package repository

import (
	"banking-system/config"
	"banking-system/models"
)

type LoanRepository struct {
}

func NewLoanRepository() *LoanRepository {
	return &LoanRepository{}
}

func (r *LoanRepository) Create(loan *models.Loan) error {
	return config.DB.Create(loan).Error
}

func (r *LoanRepository) GetAll() ([]models.Loan, error) {

	var loans []models.Loan

	err := config.DB.Find(&loans).Error

	return loans, err
}
