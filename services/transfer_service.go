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
get the recipient's account by account number
to not transfer to self
check if user has sufficient balance
update account in the database for both sender and recipient
create a transaction record for the transfer
*/

func (s *TransferServices) TransferFunds(senderUserID uint, request *models.TransferRequest) error {

	// Get logged in user account which is the sender account

	senderAcc, err := s.Repo.GetAccountByUserID(senderUserID)
	if err != nil {
		return errors.New("sender account not found")
	}

	// recipient account by account number

	recipientAcc, err := s.Repo.GetAccountByAccountNumber(request.RecipientAcc)
	if err != nil {
		return errors.New("recipient account not found")
	}

	// not allow transfer to self

	if senderAcc.AccountNumber == recipientAcc.AccountNumber {
		return errors.New("cannot transfer to your own account")
	}

	// check users balance, if user balance is less than the amount they want to transfer error here
	if senderAcc.Balance < request.Amount {
		return errors.New("insufficient balance")
	}

	// deduct amount from sender's balance and add to recipient's balance
	fmt.Printf("Transferring %.2f from %s to %s\n", request.Amount, senderAcc.AccountNumber, recipientAcc.AccountNumber)
	senderAcc.Balance -= request.Amount
	recipientAcc.Balance += request.Amount

	// Save updated sender account balance
	err = s.Repo.UpdateAccount(senderAcc)
	if err != nil {
		return errors.New("failed to update sender account balance")
	}

	// Save updated recipient account balance
	err = s.Repo.UpdateAccount(recipientAcc)
	if err != nil {
		return errors.New("failed to update recipient account balance")
	}

	// Create transfer transaction record
	err = s.Repo.CreateTransfer(senderAcc.ID, recipientAcc.ID, request.Amount, request.Description)
	if err != nil {
		return fmt.Errorf("failed to create transfer transaction: %w", err)
	}

	return nil
}
