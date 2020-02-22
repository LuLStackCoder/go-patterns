package models

type account interface {
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
	Info() string
}

// Accounts represents type storing accounts
type Accounts map[uint64]account
