package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserID        uint     `json:"-" gorm:"not null"` // Foreign key to User
	AccountNumber uint64   `json:"account_number" gorm:"uniqueIndex;not null;size:20"` // Changed to string
	Balance       float64  `json:"balance" gorm:"not null;default:0"`
}

type TransferRequest struct {
	AccountNumber uint64  `json:"account_number" binding:"required"`  //changed this to match the Account model foreign key
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Description   string  `json:"description"`
}


//the user ID is associated with the transfer request and authenticated route
// figma image requests for recipient not recipient's acc num
// for every transfer, 2 transactions occur.
// 1. debit for the user
// 2. credit for the receiver
