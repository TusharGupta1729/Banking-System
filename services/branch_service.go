package services

import (
	"errors"

	"banking-system/models"
	"banking-system/repository"
)

type BranchService struct {
	repo *repository.BranchRepository
}

func NewBranchService(repo *repository.BranchRepository) *BranchService {
	return &BranchService{
		repo: repo,
	}
}

func (s *BranchService) CreateBranch(branch *models.Branch) error {

	if branch.Name == "" {
		return errors.New("branch name is required")
	}

	if branch.BankID == 0 {
		return errors.New("bank id is required")
	}

	return s.repo.Create(branch)
}

func (s *BranchService) GetBranches() ([]models.Branch, error) {
	return s.repo.GetAll()
}
