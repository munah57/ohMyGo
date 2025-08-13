package repository

import (
	"errors"
	"vaqua/db"
	"vaqua/models"

	"gorm.io/gorm"
)

type TransferRepository interface {
	GetAccountByUserID(userID uint) (*models.Account, error)
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



// we need to update the account after the transfer so it can reflect the new balance, it gets saved to the database
func (r *TransferRepo) UpdateAccount(account *models.Account) error {
	return db.Db.Save(account).Error
}

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

//account account 
func (r *TransferRepo) CreateAccount(account *models.Account) error {
	return db.Db.Create(account).Error
}

