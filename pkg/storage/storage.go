package storage

import (
	`encoding/json`
	`fmt`

	`github.com/LuLStackCoder/go-patterns/pkg/models`
)

// Storage implements possibility to work with accounts
type Storage interface {
	AddToBalance(accountID, amount uint64) error
	SubFromBalance(accountID, amount uint64) error
	Jsonify(accountID uint64) ([]byte, error)
}

type storage struct {
	accounts map[uint64]models.Account
}

// AddToBalance adds the amount to the account balance
func (s *storage) AddToBalance(accountID, amount uint64) (errGet error) {
	reqAccount, errGet := s.getAccount(accountID)
	if errGet != nil {
		return fmt.Errorf("id doesn't exist in account storage")
	}
	reqAccount.Balance += amount
	fmt.Printf("%d successfully added on wallet %s id: %d\n", amount, reqAccount.Name, reqAccount.AccountID)
	return
}

// SubFromBalance subtracts the amount from the account balance
func (s *storage) SubFromBalance(accountID, amount uint64) (errGet error) {
	reqAccount, errGet := s.getAccount(accountID)
	if errGet != nil {
		return fmt.Errorf("id doesn't exist in account storage")
	}
	if reqAccount.Balance < amount {
		return fmt.Errorf("not enough money on wallet")
	}
	reqAccount.Balance -= amount
	fmt.Printf("%d successfully debited from wallet %s id: %d\n", amount, reqAccount.Name, reqAccount.AccountID)
	return
}

// Return Info in json-[]byte
func (s *storage) Jsonify(accountID uint64) ([]byte, error) {
	var reqAccount, errGet = s.getAccount(accountID)
	if errGet != nil {
		return nil, fmt.Errorf("id doesn't exist in account storage")
	}
	var accInfo, errMarsh = json.Marshal(reqAccount)
	if errMarsh != nil {
		return nil, errMarsh
	}
	return accInfo, nil
}

// getAccount returns account instance by id
func (s *storage) getAccount(accountID uint64) (models.Account, error) {
	var reqAccount, ok = s.accounts[accountID]
	if !ok {
		return reqAccount, fmt.Errorf("id doesn't exist in account storage")
	}
	return reqAccount, nil
}

// NewStorage initializes the storage
func NewStorage(accounts map[uint64]models.Account) Storage {
	return &storage{
		accounts: accounts,
	}
}
