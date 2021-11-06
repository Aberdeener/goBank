package accounts

import "tadhg.sh/goBank/customer"

type SavingsAccount struct {
	account *Account
}

func (savingsAccount SavingsAccount) New(owner customer.Customer, interestRate float32) SavingsAccount {

	account := new(Account)
	account.New(owner, savingsAccount.Category(), interestRate)

	savingsAccount.account = account

	return savingsAccount
}

func (savingsAccount SavingsAccount) Amount() float32 {
	return savingsAccount.account.Amount
}

func (savingsAccount SavingsAccount) Owner() customer.Customer {
	return savingsAccount.account.Owner
}

func (savingsAccount SavingsAccount) Category() string {
	return "savings"
}

func (savingsAccount SavingsAccount) Deposit(amount float32) {
	savingsAccount.account.deposit(amount)
}

func (savingsAccount SavingsAccount) Withdraw(amount float32) {
	savingsAccount.account.withdraw(amount)
}

func (savingsAccount SavingsAccount) Transfer(otherAccount IAccount, amount float32) {
	savingsAccount.account.transfer(otherAccount, amount)
}

func (savingsAccount SavingsAccount) ComputeInterest() {
	savingsAccount.account.Amount = savingsAccount.account.Amount * (100 + savingsAccount.account.InterestRate) / 100
}
