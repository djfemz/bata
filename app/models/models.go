package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	AccountNumber string
	Balance       float64
}

type Transaction struct {
	gorm.Model
	AccountNumber string
	Reference     string
	Amount        float64
	Status        string
	Type          string
}
