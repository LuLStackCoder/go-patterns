package facade

type storage interface {
	AddToBalance(accountID, amount uint64) error
	SubFromBalance(accountID, amount uint64) error
	Jsonify(accountID uint64) ([]byte, error)
}

type validator interface {
	CheckCreditAmount(amount uint64) error
	CheckDebitAmount(amount uint64) error
}

// Facade represent interface to credit/debit the specific account balance by id
type Facade interface {
	Credit(accountID uint64, amount uint64) error
	Debit(accountID uint64, amount uint64) error
	GetInfo(accountID uint64) (string, error)
}

type facade struct {
	storage   storage
	validator validator
}

// Credit adds the amount to the certain account balance obtained from storage by id
func (f *facade) Credit(accountID, amount uint64) error {
	if errCheck := f.validator.CheckCreditAmount(amount); errCheck != nil {
		return errCheck
	}
	if errAdd := f.storage.AddToBalance(accountID, amount); errAdd != nil {
		return errAdd
	}
	return nil
}

// Debit remove the amount to the certain account balance obtained from storage by id
func (f *facade) Debit(accountID, amount uint64) error {
	if errCheck := f.validator.CheckDebitAmount(amount); errCheck != nil {
		return errCheck
	}
	if errSub := f.storage.SubFromBalance(accountID, amount); errSub != nil {
		return errSub
	}
	return nil
}

// Get info about the certain accountID
func (f *facade) GetInfo(accountID uint64) (string, error) {
	var accInfo, errInfo = f.storage.Jsonify(accountID)
	if errInfo != nil {
		return "", errInfo
	}
	return string(accInfo), nil
}

// NewFacade initializes the Facade
func NewFacade(storage storage, validator validator) *facade {
	return &facade{
		storage:   storage,
		validator: validator,
	}
}
