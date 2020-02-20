package account

import (
	`github.com/stretchr/testify/mock`
)

type Mock struct {
	mock.Mock
}

func (m *Mock) AddToBalance(amount uint64) error {
	args := m.Called(amount)
	return args.Get(0).(error)
}

func (m *Mock) SubFromBalance(amount uint64) error {
	args := m.Called(amount)
	return args.Get(0).(error)
}

func (m *Mock) Info() string {
	args := m.Called()
	return args.Get(0).(string)
}
