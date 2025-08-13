package repository

import (
	
	"vaqua/db"
	"vaqua/models"
	
)

type UserRepository interface {
	GetNewUserByEmail(email string) (*models.User, error)
	CreateNewUser(user *models.User) error
	CheckAccNumExists(accountNumber string) (bool, error)
}

type UserRepo struct{}

func (r *UserRepo) GetNewUserByEmail(email string) (*models.User, error) {
	var user models.User
	
	err := db.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}


func (r *UserRepo) CheckAccNumExists(accNum string) (bool, error) {
    var count int64
    err := db.Db.Model(&models.User{}).Where("account_number = ?", accNum).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}


func (r *UserRepo) CreateNewUser(user *models.User) error {
	err := db.Db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
	
}


