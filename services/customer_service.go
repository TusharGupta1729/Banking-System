package services

import (
	"errors"

	"banking-system/models"
	"banking-system/repository"
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

	return s.repo.Create(customer)
}

func (s *CustomerService) GetCustomers() ([]models.Customer, error) {
	return s.repo.GetAll()
}
