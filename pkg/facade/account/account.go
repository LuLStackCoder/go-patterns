package account

import (
	"fmt"
)

type Account interface {
	Info() string
	ValidateName(accountName string) error
	ValidatePhone(accountPhone string) error
}

type account struct {
	name  string
	phone string
}

func (a *account) Info() string {
	return fmt.Sprintf("{\n\t\"name\": \"%s\",\n\t\"phone\": \"%s\"\n}", a.name, a.phone)
}

func (a *account) ValidateName(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("account name verified")
	return nil
}

func (a *account) ValidatePhone(accountPhone string) error {
	if a.phone != accountPhone {
		return fmt.Errorf("account phone is invalid")
	}
	fmt.Println("account name valid")
	return nil
}

func NewAccount(name, phone string) *account {
	return &account{name: name, phone: phone}
}
