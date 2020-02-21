package account

import (
	`testing`

	`github.com/stretchr/testify/assert`
)

const (
	expectedStringAlex = `{
	"accountID": 1,
	"name": "AlexMercer",
	"cardNumber": "427623452142",
	"cvv": "772",
	"balance": 1200
}`
)

// var arrFields = []fields{
// 	{0, "JamesBond", "427623499434", "221", 450},
// 	{1, "AlexMercer", "427623452142", "772", 1200},
// 	{2, "EdsgerDijkstra", "427621234151", "355", 3400},
// 	{3, "AlanTuring", "42762948753743", "987", 5000},
// }

type fields struct {
	accountID  uint64
	name       string
	cardNumber string
	cvv        string
	balance    uint64
}

func TestAccountAddToBalance(t *testing.T) {
	type args struct {
		amount uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"TestErrAddMinLim", args{0}, true},
		{"TestErrAddMaxLim", args{50000}, true},
		{"TestAddOk", args{1000}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAccount(1, "AlexMercer", "427623452142", "772", 1200)
			if err := a.AddToBalance(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("AddToBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountSubFromBalance(t *testing.T) {
	type args struct {
		amount uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"TestErrSub", args{2200}, true},
		{"TestOkSub", args{100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAccount(1, "AlexMercer", "427623452142", "772", 1200)
			if err := a.SubFromBalance(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SubFromBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountInfo(t *testing.T) {
	tests := []struct {
		name   string
		want   string
	}{
		{"TestInfoAlex", expectedStringAlex},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAccount(1, "AlexMercer", "427623452142", "772", 1200)
			assert.Equal(t, tt.want, a.Info())
		})
	}
}
