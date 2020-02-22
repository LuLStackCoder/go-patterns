package main

import (
	`fmt`
	`os`

	`github.com/LuLStackCoder/go-patterns/pkg/account`
	`github.com/LuLStackCoder/go-patterns/pkg/facade`
	`github.com/LuLStackCoder/go-patterns/pkg/models`
)

var (
	accounts = models.Accounts{
		0: account.NewAccount(0, "JamesBond", "427623499434", "221", 450),
		1: account.NewAccount(1, "AlexMercer", "427623452142", "772", 1200),
		2: account.NewAccount(2, "EdsgerDijkstra", "427621234151", "355", 3400),
		3: account.NewAccount(3, "AlanTuring", "42762948753743", "987", 5000),
	}
)

func main() {
	var paymentSystem = facade.NewPayment(accounts)
	var infoId0, errId0 = paymentSystem.GetInfo(0)
	if errId0 == nil {
		fmt.Println(infoId0)
		if err := paymentSystem.Credit(0, 220); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		infoId0, _ = paymentSystem.GetInfo(0)
		fmt.Println(infoId0)
	}
	var infoId2, errId2 = paymentSystem.GetInfo(2)
	if errId2 == nil {
		fmt.Println(infoId2)
		if err := paymentSystem.Debit(2, 3000); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		infoId2, _ = paymentSystem.GetInfo(2)
		fmt.Println(infoId2)
	}
	var infoId1, errId1 = paymentSystem.GetInfo(1)
	if errId1 == nil {
		fmt.Println(infoId1)
	}
}
