package facade

import (
	`fmt`
)

type storageFacade interface {
	AddToBalance(accountID, amount uint64) error
	SubFromBalance(accountID, amount uint64) error
	Jsonify(accountID uint64) ([]byte, error)
}

type validatorFacade interface {
	CheckCreditAmount(amount uint64) bool
	CheckDebitAmount(amount uint64) bool
}

// Facade represent interface to credit/debit the specific account balance by id
type Facade interface {
	Credit(accountID uint64, amount uint64) error
	Debit(accountID uint64, amount uint64) error
	GetInfo(accountID uint64) (string, error)
}

type facade struct {
	storage   storageFacade
	validator validatorFacade
}

// Credit adds the amount to the certain account balance obtained from storage by id
func (f *facade) Credit(accountID, amount uint64) (err error) {
	if ok := f.validator.CheckCreditAmount(amount); !ok {
		return fmt.Errorf("amount exceeds the credit limit")
	}
	return f.storage.AddToBalance(accountID, amount)
}

// Debit remove the amount to the certain account balance obtained from storage by id
func (f *facade) Debit(accountID, amount uint64) (err error) {
	if ok := f.validator.CheckDebitAmount(amount); !ok {
		return fmt.Errorf("amount exceeds the debit limit")
	}
	return f.storage.SubFromBalance(accountID, amount)
}

// Get info about the certain accountID
func (f *facade) GetInfo(accountID uint64) (accInfo string, err error) {
	var byteInfo []byte
	if byteInfo, err = f.storage.Jsonify(accountID); err != nil {
		return "", err
	}
	accInfo = string(byteInfo)
	return
}

// NewFacade initializes the Facade
func NewFacade(storage storageFacade, validator validatorFacade) Facade {
	return &facade{
		storage:   storage,
		validator: validator,
	}
}
