package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	AccountID uint

	Type   string
	Amount float64
}
