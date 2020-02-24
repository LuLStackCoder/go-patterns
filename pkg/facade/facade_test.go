package facade

import (
	`encoding/json`
	`fmt`
	`testing`

	`github.com/stretchr/testify/assert`

	`github.com/LuLStackCoder/go-patterns/pkg/models`
	storage2 `github.com/LuLStackCoder/go-patterns/pkg/storage`
	validator2 `github.com/LuLStackCoder/go-patterns/pkg/validator`
)

const (
	methodJsonify = "Jsonify"
	methodAddToBalance   = "AddToBalance"
	methodSubFromBalance = "SubFromBalance"
	methodCheckCreditAmount = "CheckCreditAmount"
	methodCheckDebitAmount = "CheckDebitAmount"

	successID uint64 = 1
	failID    uint64 = 32

	successCreditAmount uint64 = 3000
	failCreditAmount    uint64 = 1

	successDebitAmount uint64 = 100
	failDebitAmount    uint64 = 10000000
	failDebitAcc1Amount uint64 = 5000
)

var (
	badCreditError  = fmt.Errorf("amount exceeds the credit limit")
	badDebitError = fmt.Errorf("amount exceeds the debit limit")
	badSubError  = fmt.Errorf("not enough money on wallet")
	badIdError   = fmt.Errorf("id doesn't exist in account storage")
	accounts = map[uint64]models.Account{
		0: models.Account{AccountID: 0, Name: "JamesBond", CardNumber: "427623499434", Cvv: "221", Balance: 450},
		1: models.Account{AccountID: 1, Name: "AlexMercer", CardNumber: "427623452142", Cvv: "772", Balance: 1200},
		2: models.Account{AccountID: 2, Name: "EdsgerDijkstra", CardNumber: "427621234151", Cvv: "355", Balance: 3400},
		3: models.Account{AccountID: 3, Name: "AlanTuring", CardNumber: "42762948753743", Cvv: "987", Balance: 5000},
	}
	infoId1, _ = json.Marshal(accounts[1])
)

func TestFacadeGetInfo(t *testing.T) {
	type args struct {
		accountID uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "TestInfoOk",
			args:    args{1},
			want:    string(infoId1),
			wantErr: false,
		},
		{
			name:    "TestBadId",
			args:    args{32},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := new(storage2.Mock)
			storageMock.On(methodJsonify, successID).Return(infoId1, nil).Once()
			storageMock.On(methodJsonify, failID).Return(nil, badIdError).Once()

			validatorMock := new(validator2.Mock)
			validatorMock.On(methodCheckCreditAmount, successID).Return(infoId1, nil).Once()
			validatorMock.On(methodCheckDebitAmount, failID).Return(nil, badIdError).Once()

			f := NewFacade(storageMock, validatorMock)
			got, err := f.GetInfo(tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFacadeCredit(t *testing.T) {
	type args struct {
		accountID uint64
		amount    uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestCreditGoodId",
			args:    args{successID, successCreditAmount},
			wantErr: false,
		},
		{
			name:    "TestCreditBadId",
			args:    args{failID, successCreditAmount},
			wantErr: true,
		},
		{
			name:    "TestCreditBadAmount",
			args:    args{successID, failCreditAmount},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := new(storage2.Mock)
			storageMock.On(methodAddToBalance, successID, successCreditAmount).Return(nil).Once()
			storageMock.On(methodAddToBalance, failID, successCreditAmount).Return(badIdError).Once()

			validatorMock := new(validator2.Mock)
			validatorMock.On(methodCheckCreditAmount, successCreditAmount).Return(nil).Once()
			validatorMock.On(methodCheckCreditAmount, failCreditAmount).Return(badCreditError).Once()

			p := NewFacade(storageMock, validatorMock)
			if err := p.Credit(tt.args.accountID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Credit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPaymentDebit(t *testing.T) {
	type args struct {
		accountID uint64
		amount    uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestDebitGoodId",
			args:    args{successID, successDebitAmount},
			wantErr: false,
		},
		{
			name:    "TestDebitBadId",
			args:    args{failID, successDebitAmount},
			wantErr: true,
		},
		{
			name:    "TestDebitBadAmount",
			args:    args{successID, failDebitAmount},
			wantErr: true,
		},
		{
			name:    "TestDebitBadAcc1Amount",
			args:    args{successID, failDebitAcc1Amount},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := new(storage2.Mock)
			storageMock.On(methodSubFromBalance, successID, successDebitAmount).Return(nil).Once()
			storageMock.On(methodSubFromBalance, failID, successDebitAmount).Return(badIdError).Once()
			storageMock.On(methodSubFromBalance, successID, failDebitAcc1Amount).Return(badSubError).Once()

			validatorMock := new(validator2.Mock)
			validatorMock.On(methodCheckDebitAmount, successDebitAmount).Return(nil).Once()
			validatorMock.On(methodCheckDebitAmount, failDebitAcc1Amount).Return(nil).Once()
			validatorMock.On(methodCheckDebitAmount, failDebitAmount).Return(badDebitError).Once()

			p := NewFacade(storageMock, validatorMock)
			if err := p.Debit(tt.args.accountID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Debit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
