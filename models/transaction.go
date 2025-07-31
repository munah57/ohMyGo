package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID      uint    `json:"user_id" gorm:"not null"`
	RecipientID uint    `json:"recipient_id" gorm:"not null"`
	Type        string  `json:"type" gorm:"not null"`
	Amount      float64 `json:"amount" gorm:"not null"`
	Description string  `json:"description" validate:"required,max=50"` //payment,subscription,transfer,deposit
}


