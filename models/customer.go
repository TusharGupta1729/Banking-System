package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model

	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Phone        string `gorm:"unique"`
	PasswordHash string `gorm:"not null"`
}
