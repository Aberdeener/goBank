package accounts

import "tadhg.sh/goBank/customer"

type CreditAccount struct {
	account *Account
}

func (creditAccount CreditAccount) New(owner customer.Customer, interestRate float32, isCreditCard bool) CreditAccount {

	account := new(Account)

	if isCreditCard {
		account.New(owner, "creditCard", 20)
	} else {
		account.New(owner, creditAccount.Category(), interestRate)
	}

	creditAccount.account = account

	return creditAccount
}

func (creditAccount CreditAccount) Amount() float32 {
	return creditAccount.account.Amount
}

func (creditAccount CreditAccount) Owner() customer.Customer {
	return creditAccount.account.Owner
}

func (creditAccount CreditAccount) Category() string {
	return "credit"
}

func (creditAccount CreditAccount) Deposit(amount float32) {
	creditAccount.account.deposit(amount)
}

func (creditAccount CreditAccount) Withdraw(amount float32) {
	creditAccount.account.withdraw(amount)
}

func (creditAccount CreditAccount) Transfer(otherAccount IAccount, amount float32) {
	creditAccount.account.transfer(otherAccount, amount)
}

func (creditAccount CreditAccount) ComputeInterest() {
	if creditAccount.account.Amount < 0 {
		creditAccount.account.Amount = (creditAccount.account.Amount - 10.00) * (100 + creditAccount.account.InterestRate) / 100
	}
}
