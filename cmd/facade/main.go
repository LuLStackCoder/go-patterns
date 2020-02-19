package main

import (
	`fmt`
	`os`

	`github.com/LuLStackCoder/go-patterns/pkg/facade/account`
	`github.com/LuLStackCoder/go-patterns/pkg/facade/payment`
	`github.com/LuLStackCoder/go-patterns/pkg/facade/storage`
)

func main() {
	var clients [4]account.Account
	clients[0] = account.NewAccount(0, "JamesBond", "427623499434", "221", 450)
	clients[1] = account.NewAccount(1, "AlexMercer", "427623452142", "772", 1200)
	clients[2] = account.NewAccount(2, "EdsgerDijkstra", "427621234151", "355", 3400)
	clients[3] = account.NewAccount(3, "AlanTuring", "42762948753743", "987", 5000)
	var accounts = make(map[uint64]account.Account)
	for i, _ := range clients {
		accounts[uint64(i)] = clients[i]
	}
	for i, _ := range accounts {
		fmt.Println(accounts[i])
	}
	var newStorage = storage.NewStorage(accounts)
	paymentSystem := payment.NewPayment(newStorage)
	if err := paymentSystem.AddMoney(0, "JamesBond", "427623499434", "221", 220); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if err := paymentSystem.DeductMoney(2, "EdsgerDijkstra", "427621234151", "355", 3000); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for i, _ := range accounts {
		fmt.Println(accounts[i])
	}

}
