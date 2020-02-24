package validator

import (
	`fmt`
)

// Validator implements possibility to check amount correctness
type Validator interface {
	CheckCreditAmount(amount uint64) error
	CheckDebitAmount(amount uint64) error
}

type validator struct {
	creditMaxLimit uint64
	creditMinLimit uint64
	debitMaxLimit  uint64
	debitMinLimit  uint64
}

// CheckCreditAmount check amount inside credit limits
func (v *validator) CheckCreditAmount(amount uint64) error {
	if v.creditMinLimit > amount || amount > v.creditMaxLimit {
		return fmt.Errorf("amount exceeds the credit limit")
	}
	return nil
}

// CheckDebitAmount check amount inside debit limits
func (v *validator) CheckDebitAmount(amount uint64) error {
	if v.debitMinLimit > amount || amount > v.debitMaxLimit {
		return fmt.Errorf("amount exceeds the debit limit")
	}
	return nil
}

// NewValidator initializes the validator
func NewValidator(creditMaxLimit uint64, creditMinLimit uint64,
	debitMaxLimit uint64, debitMinLimit uint64) Validator {
	return &validator{
		creditMaxLimit: creditMaxLimit,
		creditMinLimit: creditMinLimit,
		debitMaxLimit:  debitMaxLimit,
		debitMinLimit:  debitMinLimit,
	}
}
