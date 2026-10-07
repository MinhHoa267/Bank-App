package services

import (
	"errors"
	"fmt"
	"myBankApp/models"
)

func UpdateAccount(updated models.Account) error {
	accounts, err := LoadAccount()
	if err != nil {
		return err
	}

	for i := range accounts {
		if accounts[i].Username == updated.Username {
			accounts[i] = updated
			return SaveAccounts(accounts)
		}
	}

	return errors.New("Account does not exist.")
}

func Deposit(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("Account is nil.")
	}
	if amount <= 0 {
		return errors.New(
			"Deposit amount must be greater than zero",
		)
	}

	updated := *acc
	updated.Balance += amount
	fmt.Println("Your new balance: %.2f\n", acc.Balance)
	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}

func Withdraw(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("account is nil")
	}
	if amount <= 0 {
		return errors.New(
			"withdraw amount must be greater than zero",
		)
	}
	if amount > acc.Balance {
		return errors.New("account balance is not enough")
	}

	updated := *acc
	updated.Balance -= amount

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}
func Transfer(from *models.Account, toUser string, amount float64) error {
	if from == nil {
		return errors.New("sender account is nil")
	}

	if amount <= 0 {
		return errors.New(
			"transfer amount must be greater than zero",
		)
	}

	if from.Username == toUser {
		return errors.New(
			"sender and receiver must be different accounts",
		)
	}

	accounts, err := LoadAccount()
	if err != nil {
		return err
	}

	var sender *models.Account
	var receiver *models.Account

	for i := range accounts {
		if accounts[i].Username == from.Username {
			sender = &accounts[i]
		}

		if accounts[i].Username == toUser {
			receiver = &accounts[i]
		}
	}

	if sender == nil || receiver == nil {
		return errors.New("Sender or Receiver account is not existed!")
	}

	if sender.Balance < amount {
		return errors.New("Not enough cash, please deposit more!")
	}

	sender.Balance -= amount
	receiver.Balance += amount

	if err := SaveAccounts(accounts); err != nil {
		return err
	}

	from.Balance = sender.Balance

	return nil
}
