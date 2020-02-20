package facade

import (
	`fmt`
	`testing`

	account2 `github.com/LuLStackCoder/go-patterns/pkg/account`
	storage2 `github.com/LuLStackCoder/go-patterns/pkg/storage`
)

const (
	methodGetAccount = "GetAccount"

	methodAddToBalance   = "AddToBalance"
	methodSubFromBalance = "SubFromBalance"
	methodInfo           = "Info"

	successID uint64 = 1
	failID    uint64 = 32

	successAddAmount uint64 = 3000
	failAddAmount    uint64 = 0

	successSubAmount uint64 = 100
	failSubAmount    uint64 = 10000

	infoRes = `{
	"accountID": 0,
	"name": "JamesBond",
	"cardNumber": "427623499434",
	"cvv": "221",
	"balance": 450
}`
)

var (
	badAddError  = fmt.Errorf("amount exceeds the replenishment limit")
	badSubError  = fmt.Errorf("not enough money on wallet")
	badIdError   = fmt.Errorf("id doesn't exist in account storage")
	accountById1 = account2.NewAccount(1, "AlexMercer", "427623452142", "772", 1200)
)

type fields struct {
	storage storage
}
type args struct {
	accountID uint64
	amount    uint64
}

func TestPaymentCredit(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestAddOkAmount",
			args:    args{successID, successAddAmount},
			wantErr: false,
		},
		{
			name:    "TestAddBadAmount",
			args:    args{successID, failAddAmount},
			wantErr: true,
		},
		{
			name:    "TestAddBadId",
			args:    args{failID, successAddAmount},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountMock := new(account2.Mock)

			accountMock.On(methodAddToBalance, successAddAmount).Return(nil).Once()
			accountMock.On(methodAddToBalance, failAddAmount).Return(badAddError).Once()

			storageMock := new(storage2.Mock)

			storageMock.On(methodGetAccount, successID).Return(accountById1, nil).Once()
			storageMock.On(methodGetAccount, failID).Return(nil, badIdError).Once()
			p := NewPayment(storageMock)
			if err := p.Credit(tt.args.accountID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Credit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPaymentDebit(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestSubOkAmount",
			args:    args{successID, successSubAmount},
			wantErr: false,
		},
		{
			name:    "TestSubBadAmount",
			args:    args{successID, failSubAmount},
			wantErr: true,
		},
		{
			name:    "TestSubBadId",
			args:    args{failID, successSubAmount},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountMock := new(account2.Mock)

			accountMock.On(methodSubFromBalance, successSubAmount).Return(nil).Once()
			accountMock.On(methodSubFromBalance, failSubAmount).Return(badSubError).Once()

			storageMock := new(storage2.Mock)

			storageMock.On(methodGetAccount, successID).Return(accountById1, nil).Once()
			storageMock.On(methodGetAccount, failID).Return(nil, badIdError).Once()

			p := NewPayment(storageMock)
			if err := p.Debit(tt.args.accountID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Debit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
