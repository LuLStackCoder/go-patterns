package validator

import (
	`github.com/stretchr/testify/mock`
)

// Mock of validator
type Mock struct {
	mock.Mock
}

// CheckCreditAmount ...
func (m *Mock) CheckCreditAmount(amount uint64) error {
	args := m.Called(amount)
	return args.Error(0)
}

// CheckDebitAmount check amount inside debit limits
func (m *Mock) CheckDebitAmount(amount uint64) error {
	args := m.Called(amount)
	return args.Error(0)
}
