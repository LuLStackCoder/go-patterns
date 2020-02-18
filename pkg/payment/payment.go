package payment

type account interface {
	Info() string
	ValidateName(accountName string) error
	ValidatePhone(accountPhone string) error
}

type wallet interface {

}


type facade struct {
	account *account
	wallet *wallet

}
