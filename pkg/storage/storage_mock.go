package storage

import (
	`github.com/stretchr/testify/mock`

	account2 `github.com/LuLStackCoder/go-patterns/pkg/account`
)

// Mock of storage
type Mock struct {
	mock.Mock
}

// CheckGetAccount ...
func (m *Mock) GetAccount(accountID uint64) (acc account2.Account, err error) {
	args := m.Called(accountID)
	if a, ok := args.Get(0).(account2.Account); ok {
		return a, args.Error(1)
	}
	return acc, args.Error(1)
}
