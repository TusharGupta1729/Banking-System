package repository

import (
	"banking-system/config"
	"banking-system/models"
)

type CustomerRepository struct {
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	return config.DB.Create(customer).Error
}

func (r *CustomerRepository) GetAll() ([]models.Customer, error) {

	var customers []models.Customer

	err := config.DB.Find(&customers).Error

	return customers, err
}
