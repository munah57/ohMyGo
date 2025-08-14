package repository

import (
	"errors"
	"vaqua/db"
	"vaqua/models"

	"gorm.io/gorm"
)


//ammend code to reflect the recipient's account number

type TransferRepository interface {
	GetAccountByUserID(userID uint) (*models.Account, error)
	GetAccountByAccountNumber(accountNumber string) (*models.Account, error)
	UpdateAccount(account *models.Account) error
}

type TransferRepo struct{}


func NewTransferRepo() *TransferRepo {
	return &TransferRepo{}
}

func (r *TransferRepo) GetAccountByUserID(userID uint) (*models.Account, error) {
	var account models.Account

	err := db.Db.Where("user_id = ?", userID).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// No account found for this user - handle accordingly
		return nil, errors.New("user account not found")
	} else if err != nil {
		// Other errors
		return nil, err
	}
	return &account, nil
}


// GetAccountByAccountNumber retrieves an account by its account number so we can find the recipient's account during a transfer
func (r *TransferRepo) GetAccountByAccountNumber(accountNumber string) (*models.Account, error) {
	var account models.Account
	err := db.Db.Where("account_number = ?", accountNumber).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("recipient account not found")
	} else if err != nil {
		return nil, err
	}
	return &account, nil
}

// we need to update the account after the transfer so it can reflect the new balance, it gets saved to the database
func (r *TransferRepo) UpdateAccount(account *models.Account) error {
	return db.Db.Save(account).Error
}

//here is the transfer record 

func (r *TransferRepo) CreateTransfer(senderAcc, recipientAcc uint, amount float64, description string) error {
	transfer := models.Transaction{
		UserID:      senderAcc,
		RecipientID: recipientAcc,
		Type:        "",
		Amount:      amount,
		Description: description,
	}
	return db.Db.Create(&transfer).Error
}

//
func (r *TransferRepo) CreateAccount(account *models.Account) error {
	return db.Db.Create(account).Error
}

