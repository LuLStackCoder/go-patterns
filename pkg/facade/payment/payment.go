package payment

type account interface {
	Info() string
	ValidateName(accountName string) error
	ValidatePhone(accountPhone string) error
}

type wallet interface {
	AddToBalance(amount int)
	SubFromBalance(amount int) error
}

type Payment interface {
	AddMoney(accountName string, amount int) error
	DeductMoney(accountName string, amount int) error
}

type payment struct {
	account account
	wallet  wallet
}

func (p *payment) AddMoney(accountName string, amount int) error {
	if err := p.account.ValidateName(accountName); err != nil {
		return err
	}
	p.wallet.AddToBalance(amount)
	return nil
}

func (p *payment) DeductMoney(accountName string, amount int) error {
	if err := p.account.ValidateName(accountName); err != nil {
		return err
	}
	if err := p.wallet.SubFromBalance(amount); err != nil {
		return err
	}
	return nil
}

func NewPayment(account account, wallet wallet) *payment {
	return &payment{account: account, wallet: wallet}
}
