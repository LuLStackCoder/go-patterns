package payment

type account = interface {
	ValidateAccount(accountName, cardNumber, cvv string) error
	Info() (string, error)
	AddToBalance(amount uint64) error
	SubFromBalance(amount uint64) error
}

type storage = interface {
	GetAccount(accountID uint64) (account, error)
}

type Payment interface {
	AddMoney(accountID uint64, accountName, cardNumber, cvv string, amount uint64) error
	DeductMoney(accountID uint64, accountName, cardNumber, cvv string, amount uint64) error
}

type payment struct {
	storage storage
}

func (p *payment) AddMoney(accountID uint64, accountName, cardNumber, cvv string, amount uint64) error {
	account, errGetAccount := p.storage.GetAccount(accountID)
	if errGetAccount != nil {
		return errGetAccount
	}
	if errValidate := account.ValidateAccount(accountName, cardNumber, cvv); errValidate != nil {
		return errValidate
	}
	if errAdd := account.AddToBalance(amount); errAdd != nil {
		return errAdd
	}
	return nil
}

func (p *payment) DeductMoney(accountID uint64, accountName, cardNumber, cvv string, amount uint64) error {
	account, errGetAccount := p.storage.GetAccount(accountID)
	if errGetAccount != nil {
		return errGetAccount
	}
	if errValidate := account.ValidateAccount(accountName, cardNumber, cvv); errValidate != nil {
		return errValidate
	}
	if errAdd := account.SubFromBalance(amount); errAdd != nil {
		return errAdd
	}
	return nil
}

func NewPayment(storage storage) Payment {
	return &payment{
		storage: storage,
	}
}
