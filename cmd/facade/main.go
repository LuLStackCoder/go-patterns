package main

import (
	`fmt`
	`os`

	`github.com/LuLStackCoder/go-patterns/pkg/facade`
	`github.com/LuLStackCoder/go-patterns/pkg/models`
	`github.com/LuLStackCoder/go-patterns/pkg/storage`
	`github.com/LuLStackCoder/go-patterns/pkg/validator`
)

const (
	CreditMaxLimit = 100000
	CreditMinLimit = 1
	DebitMaxLimit  = 50000
	DebitMinLimit  = 1
)

var (
	accounts = map[uint64]*models.Account{
		0: {AccountID: 0, Name: "JamesBond", CardNumber: "427623499434", Cvv: "221", Balance: 450},
		1: {AccountID: 1, Name: "AlexMercer", CardNumber: "427623452142", Cvv: "772", Balance: 1200},
		2: {AccountID: 2, Name: "EdsgerDijkstra", CardNumber: "427621234151", Cvv: "355", Balance: 3400},
		3: {AccountID: 3, Name: "AlanTuring", CardNumber: "42762948753743", Cvv: "987", Balance: 5000},
	}
)

func main() {
	var newStorage = storage.NewStorage(accounts)
	var newValidator = validator.NewValidator(CreditMaxLimit, CreditMinLimit, DebitMaxLimit, DebitMinLimit)
	var newFacade = facade.NewFacade(newStorage, newValidator)
	var infoId0, errId0 = newFacade.GetInfo(0)
	if errId0 == nil {
		fmt.Println(infoId0)
		if err := newFacade.Credit(0, 220); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		infoId0, _ = newFacade.GetInfo(0)
		fmt.Println(infoId0)
	}
	var infoId2, errId2 = newFacade.GetInfo(2)
	if errId2 == nil {
		fmt.Println(infoId2)
		if err := newFacade.Debit(2, 3000); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		infoId2, _ = newFacade.GetInfo(2)
		fmt.Println(infoId2)
	}
	var infoId1, errId1 = newFacade.GetInfo(1)
	if errId1 == nil {
		fmt.Println(infoId1)
	}
}
