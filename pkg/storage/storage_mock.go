package storage

import (
	account2 `github.com/LuLStackCoder/go-patterns/pkg/account`
	`github.com/stretchr/testify/mock`
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetAccount(accountID uint64) (acc account2.Account, err error) {
	args := m.Called(accountID)
	if a, ok := args.Get(0).(account2.Account); ok {
		return a, args.Error(1)
	}
	return acc, args.Error(1)

}
