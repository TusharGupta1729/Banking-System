package services

import (
	"banking-system/models"
	"banking-system/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) error {

	if customer.Name == "" {
		return errors.New("customer name is required")
	}

	if customer.Email == "" {
		return errors.New("email is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(customer.PasswordHash),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	customer.PasswordHash = string(hashedPassword)

	customer.Role = "customer"

	return s.repo.Create(customer)
}

func (s *CustomerService) GetCustomers() ([]models.Customer, error) {
	return s.repo.GetAll()
}

func (s *CustomerService) GetCustomerAccounts(id uint) ([]models.Account, error) {

	customer, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return customer.Accounts, nil
}
