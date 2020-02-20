package storage

import (
	`testing`

	account2 `github.com/LuLStackCoder/go-patterns/pkg/account`
	`github.com/stretchr/testify/assert`
)

var (
	accountMap = map[uint64]account2.Account{
		0: account2.NewAccount(0, "JamesBond", "427623499434", "221", 450),
		1: account2.NewAccount(1, "AlexMercer", "427623452142", "772", 1200),
		2: account2.NewAccount(2, "EdsgerDijkstra", "427621234151", "355", 3400),
		3: account2.NewAccount(3, "AlanTuring", "42762948753743", "987", 5000),
	}
)

func TestStorageGetAccount(t *testing.T) {
	type fields struct {
		accounts map[uint64]account
	}
	type args struct {
		accountID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantAcc account
		wantErr bool
	}{
		{
			name:    "TestGetAccOk",
			fields:  fields{accountMap},
			args:    args{0},
			wantAcc: account2.NewAccount(0, "JamesBond", "427623499434", "221", 450),
			wantErr: false,
		},
		{
			name:    "TestGetAccFail",
			fields:  fields{accountMap},
			args:    args{32},
			wantAcc: account2.NewAccount(0, "JamesBond", "427623499434", "221", 450),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newStorage := NewStorage(accountMap)
			gotAccount, gotErr := newStorage.GetAccount(tt.args.accountID)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if gotErr == nil {
				assert.Equal(t, tt.wantAcc, gotAccount)
			}
		})
	}
}
