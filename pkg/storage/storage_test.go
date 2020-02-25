package storage

import (
	`encoding/json`
	`testing`

	`github.com/stretchr/testify/assert`

	`github.com/LuLStackCoder/go-patterns/pkg/models`
)

var (
	accounts = map[uint64]*models.Account{
		0: &models.Account{AccountID: 0, Name: "JamesBond", CardNumber: "427623499434", Cvv: "221", Balance: 450},
		1: &models.Account{AccountID: 1, Name: "AlexMercer", CardNumber: "427623452142", Cvv: "772", Balance: 1200},
		2: &models.Account{AccountID: 2, Name: "EdsgerDijkstra", CardNumber: "427621234151", Cvv: "355", Balance: 3400},
		3: &models.Account{AccountID: 3, Name: "AlanTuring", CardNumber: "42762948753743", Cvv: "987", Balance: 5000},
	}
	accountsChanged = map[uint64]*models.Account{
		0: &models.Account{AccountID: 0, Name: "JamesBond", CardNumber: "427623499434", Cvv: "221", Balance: 350},
		1: &models.Account{AccountID: 1, Name: "AlexMercer", CardNumber: "427623452142", Cvv: "772", Balance: 6200},
		2: &models.Account{AccountID: 2, Name: "EdsgerDijkstra", CardNumber: "427621234151", Cvv: "355", Balance: 3400},
		3: &models.Account{AccountID: 3, Name: "AlanTuring", CardNumber: "42762948753743", Cvv: "987", Balance: 5000},
	}
)

func TestStorageJsonify(t *testing.T) {
	type args struct {
		accountID uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestStorageJsonifyOk",
			args:    args{1},
			wantErr: false,
		},
		{
			name:    "TestStorageSubFromBalanceBadId",
			args:    args{32},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorage(accounts)
			got, err := s.Jsonify(tt.args.accountID)
			want, _ := json.Marshal(accounts[tt.args.accountID])
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, want, got)
			}
		})
	}
}

func TestStorageAddToBalance(t *testing.T) {
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
			name:    "TestStorageAddToBalanceOk",
			args:    args{1, 5000},
			wantErr: false,
		},
		{
			name:    "TestStorageAddToBalanceBadId",
			args:    args{32, 2000},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorage(accounts)
			err := s.AddToBalance(tt.args.accountID, tt.args.amount)
			assert.Equal(t, tt.wantErr, err != nil)
			want := accountsChanged[tt.args.accountID]
			got := accounts[tt.args.accountID]
			assert.Equal(t, want, got)
		})
	}
}

func TestStorageSubFromBalance(t *testing.T) {
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
			name:    "TestStorageSubFromBalanceOk",
			args:    args{0, 100},
			wantErr: false,
		},
		{
			name:    "TestStorageSubFromBalanceBadId",
			args:    args{32, 300},
			wantErr: true,
		},
		{
			name:    "TestStorageSubFromBalanceBadAmount",
			args:    args{0, 2000},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorage(accounts)
			err := s.SubFromBalance(tt.args.accountID, tt.args.amount)
			assert.Equal(t, tt.wantErr, err != nil)
			want := accountsChanged[tt.args.accountID]
			got := accounts[tt.args.accountID]
			assert.Equal(t, want, got)
		})
	}
}
