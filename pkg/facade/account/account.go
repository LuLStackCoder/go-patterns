package account

import (
	"encoding/json"
	"fmt"
)

const (
	ReplenishmentMaxLimit = 30000
	ReplenishmentMinLimit = 1
)

type Account = interface {
	ValidateAccount(accountName, cardNumber, cvv string) error
	Info() (string, error)
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
}

type account struct {
	accountID  uint64
	name       string
	cardNumber string
	cvv        string
	balance    uint64
}

func (a *account) ValidateAccount(accountName, cardNumber, cvv string) error {
	if a.name != accountName || a.cardNumber != cardNumber || a.cvv != cvv {
		return fmt.Errorf("incorrect account details")
	}
	fmt.Println("correct account details")
	return nil
}

func (a *account) Info() (string, error) {
	res, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (a *account) AddToBalance(amount uint64) error {
	if amount < ReplenishmentMinLimit || amount > ReplenishmentMaxLimit {
		return fmt.Errorf("amount exceeds the replenishment limit")
	}
	a.balance += amount
	fmt.Println("successfully added on wallet")
	return nil
}

func (a *account) SubFromBalance(amount uint64) error {
	if a.balance < amount {
		return fmt.Errorf("not enough money on wallet")
	}
	a.balance -= amount
	fmt.Println("successfully debited from wallet")
	return nil
}

func NewAccount(accountID uint64, name string, cardNumber string, cvv string,
													balance uint64) Account {
	return &account{
		accountID:  accountID,
		name:       name,
		cardNumber: cardNumber,
		cvv:        cvv,
		balance:    balance,
	}
}
