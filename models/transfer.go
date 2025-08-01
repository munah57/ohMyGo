package models

type TransferRequest struct {
	RecipientAccNum float64 `json:"recipient_acc_num" validate:"required,max=10"`
	Amount          float64 `json:"amount" gorm:"not null"`
	Description     string  `json:"description"`
}

// figma image requests for recipient not recipient's acc num
// for every transfer, 2 transactions occur.
// 1. debit for the user
// 2. credit for the receiver
