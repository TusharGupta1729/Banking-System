package repository

import (
	"banking-system/config"
	"banking-system/models"
)

type BranchRepository struct {
}

func NewBranchRepository() *BranchRepository {
	return &BranchRepository{}
}

func (r *BranchRepository) Create(branch *models.Branch) error {
	return config.DB.Create(branch).Error
}

func (r *BranchRepository) GetAll() ([]models.Branch, error) {
	var branches []models.Branch

	err := config.DB.Find(&branches).Error

	return branches, err
}
