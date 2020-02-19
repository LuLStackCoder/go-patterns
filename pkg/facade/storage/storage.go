package storage

import (
	`fmt`
)

type account = interface {
	ValidateAccount(accountName, cardNumber, cvv string) error
	Info() (string, error)
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
}

type Storage = interface {
	GetAccount(accountID uint64) (account, error)
}

type storage struct {
	accounts map[uint64]account
}

func (s *storage) GetAccount(accountID uint64) (account, error) {
	reqAccount, ok := s.accounts[accountID]
	if ok == false {
		return nil, fmt.Errorf("accountID id #{accountID} don't exist in account storage")
	}
	return reqAccount, nil
}

func NewStorage(accounts map[uint64]account) Storage {
	return &storage{
		accounts: accounts,
	}
}
