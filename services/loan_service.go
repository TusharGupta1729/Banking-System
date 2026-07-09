package services

import (
	"errors"

	"banking-system/models"
	"banking-system/repository"
)

type LoanService struct {
	repo *repository.LoanRepository
}

func NewLoanService(repo *repository.LoanRepository) *LoanService {
	return &LoanService{
		repo: repo,
	}
}

func (s *LoanService) CreateLoan(loan *models.Loan) error {

	if loan.CustomerID == 0 {
		return errors.New("customer id is required")
	}

	if loan.PrincipalAmount <= 0 {
		return errors.New("principal amount must be greater than 0")
	}

	return s.repo.Create(loan)
}

func (s *LoanService) GetLoans() ([]models.Loan, error) {
	return s.repo.GetAll()
}
