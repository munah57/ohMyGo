package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname      string  `json:"firstname" gorm:"not null"`
	Lastname       string  `json:"lastname" gorm:"not null"`
	Email          string  `json:"email" gorm:"unique;not null"`
	Password       string  `json:"password" gorm:"not null"`
	AccountNumber  uint    `json:"account_number" gorm:"unique;not null"`
	Phonenumber    uint    `json:"phone_number"`
	CurrentBalance float64 `json:"current_balance"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
