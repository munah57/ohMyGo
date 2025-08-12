package repository

import (

	"vaqua/db"
	"vaqua/models"


)

type TransferRepository interface {

	GetAccountByUserID(userID uint) (*models.Account, error)
	GetAccountByAccNum(accNum string) (*models.Account, error)
	UpdateAccount(account *models.Account) error

}

type TransferRepo struct{}


 // initialize the repo for the dependency injection


func (r *TransferRepo) GetAccountByUserID(userID uint) (*models.Account, error) {
	var account models.Account
	if err := db.Db.Where("user_id = ?", userID).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *TransferRepo) GetAccountByAccNum(accNum string) (*models.Account, error) {
	var account models.Account
	if err := db.Db.Where("account_number = ?", accNum).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

 //we need to updat the account after the transfer so it can reflect the new balance, it gets saved to the database
func (r *TransferRepo) UpdateAccount(account *models.Account) error {
	return db.Db.Save(account).Error
} 