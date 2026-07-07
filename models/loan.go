package models

import "gorm.io/gorm"

type Loan struct {
	gorm.Model

	CustomerID      uint
	PrincipalAmount float64
	InterestRate    float64
	TotalAmount     float64
	PendingAmount   float64
	Status          string
}
