package account

import (
	"fmt"
)

const (
	ReplenishmentMaxLimit = 30000
	ReplenishmentMinLimit = 1
)

// Account represent an interface to change the account balance
type Account interface {
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
	Info() string
}

type account struct {
	accountID  uint64
	name       string
	cardNumber string
	cvv        string
	balance    uint64
}

// AddToBalance adds the amount to the account balance
func (a *account) AddToBalance(amount uint64) error {
	if ReplenishmentMinLimit > amount || amount > ReplenishmentMaxLimit {
		return fmt.Errorf("amount exceeds the replenishment limit")
	}
	a.balance += amount
	fmt.Printf("%d successfully added on wallet %s id: %d\n", amount, a.name, a.accountID)
	return nil
}

// SubFromBalance subtracts the amount from the account balance
func (a *account) SubFromBalance(amount uint64) error {
	if a.balance < amount {
		return fmt.Errorf("not enough money on wallet")
	}
	a.balance -= amount
	fmt.Printf("%d successfully debit from wallet %s id: %d\n", amount, a.name, a.accountID)
	return nil
}

// Return Info in json-string
func (a *account) Info() string {
	return fmt.Sprintf(`{
	"accountID": %d,
	"name": "%s",
	"cardNumber": "%s",
	"cvv": "%s",
	"balance": %d
}`, a.accountID, a.name, a.cardNumber, a.cvv, a.balance)
}

// NewAccount initializes the Account
func NewAccount(accountID uint64, name, cardNumber, cvv string, balance uint64) Account {
	return &account{
		accountID:  accountID,
		name:       name,
		cardNumber: cardNumber,
		cvv:        cvv,
		balance:    balance,
	}
}
