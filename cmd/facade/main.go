package main

import (
	"fmt"

	"github.com/LuLStackCoder/go-patterns/pkg/facade/account"
	"github.com/LuLStackCoder/go-patterns/pkg/facade/payment"
	"github.com/LuLStackCoder/go-patterns/pkg/facade/wallet"
)

func main() {
	walletClient1 := wallet.NewWallet(1, "213123132", "123")
	accountClient1 := account.NewAccount("alex", "892513243")
	paymentSystem := payment.NewPayment(accountClient1, walletClient1)
	if err := paymentSystem.AddMoney("alex", 150); err != nil {
		return
	}
	if err := paymentSystem.DeductMoney("alex", 80); err != nil {
		return
	}
	fmt.Printf("%#v\n", walletClient1)
	fmt.Printf("%#v\n", accountClient1)
	fmt.Printf("%#v\n", paymentSystem)
}
