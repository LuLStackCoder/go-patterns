package account

import (
	`testing`
	`github.com/stretchr/testify/assert`
)

const (
	expectedStringJames = `{
	"accountID": 0,
	"name": "JamesBond",
	"cardNumber": "427623499434",
	"cvv": "221",
	"balance": 450
}`
	expectedStringAlex = `{
	"accountID": 1,
	"name": "AlexMercer",
	"cardNumber": "427623452142",
	"cvv": "772",
	"balance": 1200
}`
)

type fields struct {
	accountID  uint64
	name       string
	cardNumber string
	cvv        string
	balance    uint64
}

var arrFields = []fields{{0, "JamesBond", "427623499434", "221", 450},
	{1, "AlexMercer", "427623452142", "772", 1200},
	{2, "EdsgerDijkstra", "427621234151", "355", 3400},
	{3, "AlanTuring", "42762948753743", "987", 5000},}

func Test_account_AddToBalance(t *testing.T) {
	type args struct {
		amount uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"TestCase1", arrFields[0], args{0}, true},
		{"TestCase2", arrFields[1], args{50000}, true},
		{"TestCase3", arrFields[2], args{1000}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &account{
				accountID:  tt.fields.accountID,
				name:       tt.fields.name,
				cardNumber: tt.fields.cardNumber,
				cvv:        tt.fields.cvv,
				balance:    tt.fields.balance,
			}
			if err := a.AddToBalance(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("AddToBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_account_SubFromBalance(t *testing.T) {
	type args struct {
		amount uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"TestCase1", arrFields[0], args{2200}, true},
		{"TestCase2", arrFields[1], args{100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &account{
				accountID:  tt.fields.accountID,
				name:       tt.fields.name,
				cardNumber: tt.fields.cardNumber,
				cvv:        tt.fields.cvv,
				balance:    tt.fields.balance,
			}
			if err := a.SubFromBalance(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SubFromBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_account_Info(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"TestCase1", arrFields[0], expectedStringJames},
		{"TestCase2", arrFields[1],  expectedStringAlex},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &account{
				accountID:  tt.fields.accountID,
				name:       tt.fields.name,
				cardNumber: tt.fields.cardNumber,
				cvv:        tt.fields.cvv,
				balance:    tt.fields.balance,
			}
			assert.Equal(t, a.Info(), tt.want)
		})
	}
}