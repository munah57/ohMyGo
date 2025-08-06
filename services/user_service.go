package services

import (
	"errors"
	"fmt"
	"vaqua/middleware"
	"vaqua/models"
	"vaqua/repository"
	"vaqua/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) SignUpNewUserAcct(newUser *models.SignUpRequest) error {
	fmt.Println("Service: Starting SignUpNewUserAcct...")
	_, err := s.Repo.GetNewUserByEmail(newUser.Email)
	if err == nil {
		fmt.Println("Email already in use:", newUser.Email)
		return errors.New("This email is already in use")
		
	}

	hashpass, err := utils.HashPassword(newUser.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
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
		fmt.Printf("Attempt %d: Generated Account Number = %d\n", attempts+1, accNum)

		exists, err := s.Repo.CheckAccNumExists(uint(accNum))
		if err != nil {
			fmt.Println("Error checking if account number exists:", err)
			return err
		}

		if !exists {
			user.AccountNumber = uint64(accNum)
			fmt.Println("Unique Account Number assigned:", accNum)
			break
			} else {
				fmt.Println("Account Number already exists. Retrying...")
		}
	}

	if user.AccountNumber == 0 {
		fmt.Println("Failed to generate a unique account number")
		return errors.New("A unique account number could not be generated")
	}


	fmt.Printf("Saving new user: Email = %s | AccountNumber = %d\n", user.Email, user.AccountNumber)



	err = s.Repo.CreateNewUser(user)
	if err != nil {
		fmt.Println("Error saving user:", err)
		return err
	}


	fmt.Println("User saved successfully to DB!")
	fmt.Println("Service: Finished SignUpNewUserAcct successfully")
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

