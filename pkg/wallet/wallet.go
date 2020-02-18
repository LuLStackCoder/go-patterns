package wallet

import (
	`fmt`
)

type Waller interface {
}

type wallet struct {
	balance    int
	cardNumber string
	cvv        string
}

func (w *wallet) addBalance(amount int) {
	w.balance += amount
	fmt.Printf("%s")
}

func newWallet(balance int, cardNumber string, cvv string) *wallet {
	return &wallet{balance: balance, cardNumber: cardNumber, cvv: cvv}
}
