package validator

import (
	`testing`

	`github.com/stretchr/testify/assert`
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
		want bool
	}{
		{
			name:    "TestBadMinAmount",
			args:    args{0},
			want: false,
		},
		{
			name:    "TestBadMaxAmount",
			args:    args{1000000},
			want: false,
		},
		{
			name:    "TestGoodAmount",
			args:    args{3000},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator(CreditMaxLimit, CreditMinLimit, DebitMaxLimit, DebitMinLimit)
			checkRes := v.CheckCreditAmount(tt.args.amount)
			assert.Equal(t, tt.want, checkRes)
		})
	}
}

func TestValidatorCheckDebitAmount(t *testing.T) {
	type args struct {
		amount uint64
	}
	tests := []struct {
		name    string
		args    args
		want bool
	}{
		{
			name:    "TestBadMinAmount",
			args:    args{0},
			want: false,
		},
		{
			name:    "TestBadMaxAmount",
			args:    args{1000000},
			want: false,
		},
		{
			name:    "TestGoodAmount",
			args:    args{3000},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator(CreditMaxLimit, CreditMinLimit, DebitMaxLimit, DebitMinLimit)
			checkRes := v.CheckDebitAmount(tt.args.amount)
			assert.Equal(t, tt.want, checkRes)
		})
	}
}
