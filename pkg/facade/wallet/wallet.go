package wallet

import (
	`fmt`
)

type Wallet interface {
	AddToBalance(amount int)
	SubFromBalance(amount int) error
}

type wallet struct {
	balance    int
	cardNumber string
	cvv        string
}

func (w *wallet) AddToBalance(amount int) {
	w.balance += amount
	fmt.Printf("%d succesfully added on wallet", amount)
}

func (w *wallet) SubFromBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("not enough money on wallet")
	}
	w.balance -= amount
	fmt.Printf("%d succesfully debited from wallet", amount)
	return nil
}

func NewWallet(balance int, cardNumber string, cvv string) *wallet {
	return &wallet{balance: balance, cardNumber: cardNumber, cvv: cvv}
}
