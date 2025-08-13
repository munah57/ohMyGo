package services

import (
	"errors"
	"fmt"
	"vaqua/models"
	"vaqua/repository"
)

type TransferService interface {
	TransferFunds(userID uint, req *models.TransferRequest) error
}

type TransferServices struct {
	Repo repository.TransferRepo
} //connection to transfer repo


/*
function should:
get logged in user account with is authenticated from the context 
check if user has sufficient balance
update account in the database
create a transaction record for the transfer
*/

func (s *TransferServices) TransferFunds(senderUserID uint, request *models.TransferRequest) error {


	// Get logged in user account which is the sender account

	senderAcc, err := s.Repo.GetAccountByUserID(senderUserID)
	if err != nil {
		return errors.New("sender account not found")
	}


	// check users balance, if user balance is less than the amount they want to transfer error here
	if senderAcc.Balance < request.Amount { 
		return errors.New("insufficient balance")
	}


	// deduct amount from sender's balance
    senderAcc.Balance -= request.Amount

    // Save updated sender account balance
    err = s.Repo.UpdateAccount(senderAcc)
    if err != nil {
        return errors.New("failed to update sender account balance")
    }

	//log the transfer as a transaction

	fmt.Println("Creating transfer transaction record...")
	return nil
}