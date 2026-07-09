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

	err := config.DB.
		Preload("Customer").
		Find(&loans).Error

	return loans, err
}

func (r *LoanRepository) GetByID(id uint) (*models.Loan, error) {

	var loan models.Loan

	err := config.DB.
		Preload("Customer").
		First(&loan, id).Error

	return &loan, err
}

func (r *LoanRepository) Update(loan *models.Loan) error {
	return config.DB.Save(loan).Error
}
