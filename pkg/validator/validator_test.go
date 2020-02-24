package validator

import (
	`testing`
)

const (
	CreditMaxLimit = 100000
	CreditMinLimit = 1
	DebitMaxLimit  = 50000
	DebitMinLimit  = 1
)

func TestValidatorCheckCreditAmount(t *testing.T) {
	type args struct {
		amount uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestBadMinAmount",
			args:    args{0},
			wantErr: true,
		},
		{
			name:    "TestBadMaxAmount",
			args:    args{1000000},
			wantErr: true,
		},
		{
			name:    "TestGoodAmount",
			args:    args{3000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator(CreditMaxLimit, CreditMinLimit, DebitMaxLimit, DebitMinLimit)
			if err := v.CheckCreditAmount(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CheckCreditAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validator_CheckDebitAmount(t *testing.T) {
	type args struct {
		amount uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestBadMinAmount",
			args:    args{0},
			wantErr: true,
		},
		{
			name:    "TestBadMaxAmount",
			args:    args{1000000},
			wantErr: true,
		},
		{
			name:    "TestGoodAmount",
			args:    args{3000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator(CreditMaxLimit, CreditMinLimit, DebitMaxLimit, DebitMinLimit)
			if err := v.CheckDebitAmount(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CheckDebitAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}