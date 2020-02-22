package facade

import (
	`fmt`

	`github.com/LuLStackCoder/go-patterns/pkg/models`
)

type account interface {
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
	Info() string
}

// Facade represent interface to credit/debit the specific account balance by id
type Facade interface {
	Credit(accountID uint64, amount uint64) error
	Debit(accountID uint64, amount uint64) error
	GetInfo(accountID uint64) (string, error)
}

type facade struct {
	storage models.Accounts
}

// Credit adds the amount to the certain account balance obtained from storage by id
func (f *facade) Credit(accountID uint64, amount uint64) error {
	var account, errGetAccount = f.getAccount(accountID)
	if errGetAccount != nil {
		return errGetAccount
	}
	if errAdd := account.AddToBalance(amount); errAdd != nil {
		return errAdd
	}
	return nil
}

// Debit remove the amount to the certain account balance obtained from storage by id
func (f *facade) Debit(accountID uint64, amount uint64) error {
	var account, errGetAccount = f.getAccount(accountID)
	if errGetAccount != nil {
		return errGetAccount
	}
	if errSub := account.SubFromBalance(amount); errSub != nil {
		return errSub
	}
	return nil
}

// Get info about the certain accountID
func (f *facade) GetInfo(accountID uint64) (string, error) {
	var account, errGetAccount = f.getAccount(accountID)
	if errGetAccount != nil {
		return "", errGetAccount
	}
	var jsonString = account.Info()
	return jsonString, nil
}

// getAccount returns account instance by id
func (f *facade) getAccount(accountID uint64) (account, error) {
	var reqAccount, ok = f.storage[accountID]
	if !ok {
		return nil, fmt.Errorf("id doesn't exist in account storage")
	}
	return reqAccount, nil
}

// NewPayment initializes the Facade
func NewPayment(storage models.Accounts) Facade {
	return &facade{
		storage: storage,
	}
}
