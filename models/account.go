package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	CustomerID    uint
	BranchID      uint
	AccountNumber string  `gorm:"unique;not null"`
	Balance       float64 `gorm:"default:0"`
	AccountType   string
	Status        string
}
