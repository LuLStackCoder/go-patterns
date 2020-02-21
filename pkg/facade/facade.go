package facade

type account = interface {
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
	Info() string
}

type storage = interface {
	GetAccount(accountID uint64) (account, error)
}

// Facade implements possibility to credit/debit the specific account balance by id
type Facade interface {
	Credit(accountID uint64, amount uint64) error
	Debit(accountID uint64, amount uint64) error
	GetInfo(accountID uint64) (string, error)
}

type facade struct {
	storage storage
}

// Credit adds the amount to the certain account balance obtained from storage by id
func (p *facade) Credit(accountID uint64, amount uint64) error {
	var account, errGetAccount = p.storage.GetAccount(accountID)
	if errGetAccount != nil {
		return errGetAccount
	}
	if errAdd := account.AddToBalance(amount); errAdd != nil {
		return errAdd
	}
	return nil
}

// Debit remove the amount to the certain account balance obtained from storage by id
func (p *facade) Debit(accountID uint64, amount uint64) error {
	var account, errGetAccount = p.storage.GetAccount(accountID)
	if errGetAccount != nil {
		return errGetAccount
	}
	if errSub := account.SubFromBalance(amount); errSub != nil {
		return errSub
	}
	return nil
}

// Get info about the certain accountID
func (p *facade) GetInfo(accountID uint64) (string, error) {
	var account, errGetAccount = p.storage.GetAccount(accountID)
	if errGetAccount != nil {
		return "", errGetAccount
	}
	var jsonString = account.Info()
	return jsonString, nil
}

// NewPayment initializes the Facade
func NewPayment(storage storage) Facade {
	return &facade{
		storage: storage,
	}
}
