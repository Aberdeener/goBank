package accounts

import (
	"tadhg.sh/goBank/customer"
)

type CheckingAccount struct {
	account *Account
}

func (checkingAccount CheckingAccount) New(owner customer.Customer) CheckingAccount {

	account := new(Account)
	account.New(owner, checkingAccount.Category(), 0)

	checkingAccount.account = account

	return checkingAccount
}

func (checkingAccount CheckingAccount) Amount() float32 {
	return checkingAccount.account.Amount
}

func (checkingAccount CheckingAccount) Owner() customer.Customer {
	return checkingAccount.account.Owner
}

func (checkingAccount CheckingAccount) Category() string {
	return "checking"
}

func (checkingAccount CheckingAccount) Deposit(amount float32) {
	checkingAccount.account.deposit(amount)
}

func (checkingAccount CheckingAccount) Withdraw(amount float32) {
	checkingAccount.account.withdraw(amount)
}

func (checkingAccount CheckingAccount) Transfer(otherAccount IAccount, amount float32) {
	checkingAccount.account.transfer(otherAccount, amount)
}

func (checkingAccount CheckingAccount) ComputeInterest() {

}
