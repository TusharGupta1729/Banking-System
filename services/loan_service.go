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

	loan.PendingAmount = loan.TotalAmount

	loan.Status = "Pending"
	return s.repo.Create(loan)
}

func (s *LoanService) GetLoans() ([]models.Loan, error) {
	return s.repo.GetAll()
}

func (s *LoanService) GetLoanByID(
	id uint,
) (*models.Loan, error) {

	return s.repo.GetByID(id)
}

func (s *LoanService) GetLoansByCustomerID(
	customerID uint,
) ([]models.Loan, error) {

	return s.repo.GetByCustomerID(customerID)
}

func (s *LoanService) ApproveLoan(id uint) error {

	loan, err := s.repo.GetByID(id)

	if err != nil {
		return errors.New("loan not found")
	}

	loan.Status = "Approved"

	return s.repo.Update(loan)
}

func (s *LoanService) RejectLoan(id uint) error {

	loan, err := s.repo.GetByID(id)

	if err != nil {
		return errors.New("loan not found")
	}

	loan.Status = "Rejected"

	return s.repo.Update(loan)
}

func (s *LoanService) RepayLoan(
	id uint,
	amount float64,
) error {

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	loan, err := s.repo.GetByID(id)

	if err != nil {
		return errors.New("loan not found")
	}

	if loan.Status != "Approved" {
		return errors.New("loan is not approved")
	}

	loan.PendingAmount -= amount

	if loan.PendingAmount <= 0 {
		loan.PendingAmount = 0
		loan.Status = "Closed"
	}

	return s.repo.Update(loan)
}
