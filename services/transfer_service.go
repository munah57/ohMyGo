package services

import (
	"errors"
	"vaqua/models"
	"vaqua/repository"
)

type TransferService interface {
	TransferFunds(userID uint, req models.TransferRequest) error
}

type TransferServices struct {
	Repo repository.TransferRepo
} //connection to transfer repo


/*
function should:
get logged in user account with is authenticated from the context 
get recipient account by account number
check if user has sufficient balance
check if recipient account is not the same as user's account
update both accounts in the database
create a transaction record for the transfer
*/

func (s *TransferServices) TransferFunds(userID uint, request *models.TransferRequest) error {


	// Get logged in user account which is the sender account

	sender, err := s.Repo.GetAccountByUserID(userID)
	if err != nil {
		return errors.New("account not found")
	}

	// Get recipient account by account number

	recipient, err := s.Repo.GetAccountByAccNum(request.AccountNumber)
	if err != nil {
		return errors.New("recipient account not found")
	}


	// check users balance, if user balance is less than the amount they want to transfer error here
	if sender.Balance < request.Amount { 
		return errors.New("insufficient balance")
	}


	// check if recipient account is the same as user's account
	if sender.AccountNumber == recipient.AccountNumber {
		return errors.New("cannot transfer to your own account")


	}

	//check balance of sender 
	sender.Balance = sender.Balance - request.Amount
	recipient.Balance = recipient.Balance + request.Amount

	//update both accounts in the database

	if err := s.Repo.UpdateAccount(sender); err != nil {
		return err
	}
	

	return nil
}