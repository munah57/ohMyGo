package services

import (
	"errors"
	"vaqua/middleware"
	"vaqua/models"
	"vaqua/repository"
	"vaqua/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) SignUpNewUserAcct(newUser *models.SignUpRequest) error {

		_, err := s.Repo.GetNewUserByEmail(newUser.Email)
	if err == nil {	
		return errors.New("This email is already in use")	
	}

	hashpass, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = hashpass

	user := &models.User{
		Email: newUser.Email,
		Password: hashpass,
	}

	// Generate user account number
	for attempts := 0; attempts < 5; attempts++ {
		accNum := utils.GenerateRandomAccNum()
		exists, err := s.Repo.CheckAccNumExists(uint(accNum))
		if err != nil {	
			return err
		}

		if !exists {
			user.AccountNumber = uint64(accNum)	
			break	
		}
	}

	if user.AccountNumber == 0 {	
		return errors.New("A unique account number could not be generated")
	}


	err = s.Repo.CreateNewUser(user)
	if err != nil {	
		return err
	}
	return nil
}

func (s *UserService) LoginUser(request models.LoginRequest) (string, error) {
    user, err := s.Repo.GetNewUserByEmail(request.Email)
    if err !=nil {
        return "", err
    }
    err = utils.ComparePassword(user.Password, request.Password)
    if err != nil {
        return "", err
    }
    token, err := middleware.GenerateJWT(user.ID)
    if err != nil {
        return "", err
    }
    return token, nil 
}

