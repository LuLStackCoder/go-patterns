package main

import (
	`log`
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
	logErr := log.New(os.Stderr, "", 0)
	logDef := log.New(os.Stdout, "", 0)
	newStorage := storage.NewStorage(accounts)
	newValidator := validator.NewValidator(CreditMaxLimit, CreditMinLimit, DebitMaxLimit, DebitMinLimit)
	newFacade := facade.NewFacade(newStorage, newValidator)

	if infoId0, err := newFacade.GetInfo(0); err == nil {
		logDef.Println(infoId0)
		if err := newFacade.Credit(0, 220); err != nil {
			logErr.Println(err)
		}
		infoId0, _ = newFacade.GetInfo(0)
		logDef.Println(infoId0)
	}

	if infoId2, err := newFacade.GetInfo(2); err == nil {
		logDef.Println(infoId2)
		if err := newFacade.Debit(2, 3000); err != nil {
			logErr.Println(err)
		}
		infoId2, _ = newFacade.GetInfo(2)
		logDef.Println(infoId2)
	}
}
