package accounts

import "tadhg.sh/goBank/customer"

type IAccount interface {
	Amount() float32
	Owner() customer.Customer
	Category() string
	Deposit(amount float32)
	Withdraw(amount float32)
	Transfer(otherAccount IAccount, amount float32)
	ComputeInterest()
}
