package storage

import (
	`fmt`
)

type account = interface {
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
	Info() string
}

// Storage implements possibility to get the account by id
type Storage = interface {
	GetAccount(accountID uint64) (account, error)
}

type storage struct {
	accounts map[uint64]account
}

// GetAccount returns account instance by id
func (s *storage) GetAccount(accountID uint64) (account, error) {
	var reqAccount, ok = s.accounts[accountID]
	if ok == false {
		return nil, fmt.Errorf(fmt.Sprintf("ID %d id doesn't exist in account storage", accountID))
	}
	return reqAccount, nil
}

// NewStorage initializes the storage
func NewStorage(accounts map[uint64]account) Storage {
	return &storage{
		accounts: accounts,
	}
}
