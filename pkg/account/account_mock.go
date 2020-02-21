package account

import (
	`github.com/stretchr/testify/mock`
)

// Mock of account
type Mock struct {
	mock.Mock
}

// CheckAddToBalance ...
func (m *Mock) AddToBalance(amount uint64) error {
	args := m.Called(amount)
	return args.Get(0).(error)
}

// CheckSubFromBalance ...
func (m *Mock) SubFromBalance(amount uint64) error {
	args := m.Called(amount)
	return args.Get(0).(error)
}

// CheckInfo ...
func (m *Mock) Info() string {
	args := m.Called()
	return args.Get(0).(string)
}
