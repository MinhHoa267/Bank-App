package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"myBankApp/models"
	"os"
)

const dataFile = "./data/users.json"

func LoadAccount() ([]models.Account, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("data file does not exist: %s: %w", dataFile, err)
		}
		if errors.Is(err, fs.ErrPermission) {
			return nil, fmt.Errorf("data file now allowed to read: %s: %w", dataFile, err)
		}
		return nil, fmt.Errorf("cannot read data : %s: %w", dataFile, err)
	}

	var accounts []models.Account
	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, fmt.Errorf("cannot parse data : %s: %w", dataFile, err)
	}
	return accounts, nil
}

func SaveAccounts(accounts []models.Account) error {
	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return fmt.Errorf("Cannot convert accounts : %w", err)
	}

	if err := os.WriteFile(dataFile, data, 0600); err != nil {
		return fmt.Errorf(
			"cannot write accounts  %s: %w",
			dataFile,
			err,
		)
	}

	return nil
}

func FindAccount(username string) (models.Account, error) {
	accounts, err := LoadAccount()
	if err != nil {
		return models.Account{}, err
	}
	for _, acc := range accounts {
		if acc.Username == username {
			return acc, nil
		}
	}
	return models.Account{}, fmt.Errorf("account %q does not exist", username)
}