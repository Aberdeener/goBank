package accounts

import "tadhg.sh/goBank/customer"

type CreditCardAccount struct {
	creditAccount CreditAccount
}

func (creditCardAccount CreditCardAccount) New(owner customer.Customer) CreditCardAccount {

	creditAccount := new(CreditAccount).New(owner, 0, true)

	creditCardAccount.creditAccount = creditAccount

	return creditCardAccount
}

func (creditCardAccount CreditCardAccount) Amount() float32 {
	return creditCardAccount.creditAccount.account.Amount
}

func (creditCardAccount CreditCardAccount) Owner() customer.Customer {
	return creditCardAccount.creditAccount.account.Owner
}

func (creditCardAccount CreditCardAccount) Category() string {
	return "creditCard"
}

func (creditCardAccount CreditCardAccount) Deposit(amount float32) {
	creditCardAccount.creditAccount.account.deposit(amount)
}

func (creditCardAccount CreditCardAccount) Withdraw(amount float32) {
	creditCardAccount.creditAccount.account.withdraw(amount)
}

func (creditCardAccount CreditCardAccount) Transfer(otherAccount IAccount, amount float32) {
	creditCardAccount.creditAccount.account.transfer(otherAccount, amount)
}

func (creditCardAccount CreditCardAccount) ComputeInterest() {
	if creditCardAccount.creditAccount.account.Amount < 0 {
		creditCardAccount.creditAccount.account.Amount = (creditCardAccount.creditAccount.account.Amount - 10.00) * (100 + creditCardAccount.creditAccount.account.InterestRate) / 100
	}
}
