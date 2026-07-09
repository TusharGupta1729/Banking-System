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

	err := config.DB.
		Preload("Accounts").
		Preload("Loans").
		Find(&customers).Error

	return customers, err
}

func (r *CustomerRepository) GetByEmail(email string) (*models.Customer, error) {

	var customer models.Customer

	err := config.DB.
		Where("email = ?", email).
		First(&customer).Error

	return &customer, err
}

func (r *CustomerRepository) GetByID(id uint) (*models.Customer, error) {

	var customer models.Customer

	err := config.DB.
		Preload("Accounts").
		First(&customer, id).Error

	return &customer, err
}
