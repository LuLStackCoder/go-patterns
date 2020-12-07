package validator

// Validator implements possibility to check amount correctness
type Validator interface {
	CheckCreditAmount(amount uint64) bool
	CheckDebitAmount(amount uint64) bool
}

type validator struct {
	creditMaxLimit uint64
	creditMinLimit uint64
	debitMaxLimit  uint64
	debitMinLimit  uint64
}

// CheckCreditAmount check amount inside credit limits
func (v *validator) CheckCreditAmount(amount uint64) bool {
	return v.creditMinLimit < amount && amount < v.creditMaxLimit
}

// CheckDebitAmount check amount inside debit limits
func (v *validator) CheckDebitAmount(amount uint64) bool {
	return v.debitMinLimit < amount && amount < v.debitMaxLimit
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
