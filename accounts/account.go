package accounts

import (
	"errors"
	"fmt"
	"tadhg.sh/goBank/customer"
)

type Account struct {
	Amount       float32
	Owner        customer.Customer
	Category     string
	InterestRate float32
}

func (account *Account) New(owner customer.Customer, category string, interestRate float32) *Account {
	if interestRate < 0 || interestRate > 100 {
		panic(errors.New("interestRate must be between 0 and 100"))
	}

	account.Amount = 0
	account.Owner = owner
	account.Category = category
	account.InterestRate = interestRate

	return account
}

func (account *Account) deposit(amount float32) {
	account.Amount += amount
}

func (account *Account) withdraw(amount float32) {
	account.Amount -= amount
}

func (account *Account) transfer(otherAccount IAccount, amount float32) {
	account.withdraw(amount)
	otherAccount.Deposit(amount)
}

func (account Account) getInfo() string {
	return fmt.Sprintf("Type: %s | Owner: %s | Amount: $%f\n", account.Category, account.Owner.Name, account.Amount)
}
