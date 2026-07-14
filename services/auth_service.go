package services

import (
	"banking-system/repository"
	"banking-system/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	customerRepo *repository.CustomerRepository
}

func NewAuthService(
	customerRepo *repository.CustomerRepository,
) *AuthService {
	return &AuthService{
		customerRepo: customerRepo,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthService) Login(
	req *LoginRequest,
) (string, error) {

	customer, err := s.customerRepo.GetByEmail(req.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(customer.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(
		customer.ID,
		customer.Role,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
