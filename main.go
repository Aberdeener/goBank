package main

import (
	"fmt"
	"tadhg.sh/goBank/customer"
)

func main() {
	var bank = new(Bank).New("BCIT Bank")

	var tim = new(customer.Customer).New("Tim", "A0001")
	var yuki = new(customer.Customer).New("Yuki", "A0002")
	var malik = new(customer.Customer).New("Malik", "A0003")

	bank.createAccount(yuki, "checking", 0)
	bank.createAccount(yuki, "savings", 2)

	bank.createPackage(malik)

	bank.createAccount(tim, "creditCard", 0)

	for _, account := range bank.findAccounts(yuki.Ssn, "") {
		account.Deposit(1000)
	}

	malikCheckingAccount := bank.findAccounts(malik.Ssn, "checking")[0]
	malikCheckingAccount.Deposit(1000)
	malikCreditAccount := bank.findAccounts(malik.Ssn, "credit")[0]
	malikCreditAccount.Withdraw(500)
	malikCheckingAccount.Transfer(malikCreditAccount, 500)

	timCreditCardAccount := bank.findAccounts(tim.Ssn, "creditCard")[0]
	timCreditCardAccount.Withdraw(1000)

	fmt.Printf("Total Deposits %f\n", bank.totalDeposits())
	fmt.Printf("Total Credits %f\n", bank.totalCredits())

	print("------------------------\n")

	for i := 0; i < 12; i++ {
		bank.computeInterest()
	}

	fmt.Printf("Total Deposits %f\n", bank.totalDeposits())
	fmt.Printf("Total Credits %f\n", bank.totalCredits())

	yukiSavingsAccount := bank.findAccounts(yuki.Ssn, "savings")[0]
	fmt.Printf("Yuki savings amount %f\n", yukiSavingsAccount.Amount())

	fmt.Printf("Tim credit card owing %f\n", timCreditCardAccount.Amount())
}
