package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model

	BankID   uint
	Name     string `gorm:"not null"`
	IFSCCode string `gorm:"unique;not null"`
	Address  string
}
