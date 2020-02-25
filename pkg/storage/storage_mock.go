package storage

import (
	`github.com/stretchr/testify/mock`
)

// Mock of storage
type Mock struct {
	mock.Mock
}

// AddToBalance ...
func (m *Mock) AddToBalance(accountID, amount uint64) (err error) {
	args := m.Called(accountID, amount)
	return args.Error(0)
}

// SubFromBalance ...
func (m *Mock) SubFromBalance(accountID, amount uint64) (err error) {
	args := m.Called(accountID, amount)
	return args.Error(0)
}

// Jsonify ...
func (m *Mock) Jsonify(accountID uint64) (accInfo []byte, err error) {
	args := m.Called(accountID)
	if a, ok := args.Get(0).([]byte); ok {
		return a, args.Error(1)
	}
	return accInfo, args.Error(1)
}
