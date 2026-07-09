package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	AccountID uint
	Account   Account

	Type   string
	Amount float64
}
