package models

// Account represents the account entity
type Account struct {
	AccountID  uint64 `json:"account_id"`
	Name       string `json:"name"`
	CardNumber string `json:"card_number"`
	Cvv        string `json:"cvv"`
	Balance    uint64 `json:"balance"`
}
