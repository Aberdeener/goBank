package main

import (
	"tadhg.sh/goBank/accounts"
	"tadhg.sh/goBank/customer"
)

const Checking = "checking"
const Credit = "credit"
const CreditCard = "creditCard"
const Savings = "savings"

type Bank struct {
	Name     string
	Accounts map[string][]accounts.IAccount
}

func (bank Bank) New(name string) Bank {
	bank.Name = name
	bank.Accounts = make(map[string][]accounts.IAccount)
	return bank
}

func (bank Bank) createAccount(owner customer.Customer, category string, interestRate float32) {
	ssn := owner.Ssn

	switch category {
	case Checking:
		bank.saveAccount(ssn, new(accounts.CheckingAccount).New(owner))
		break
	case Credit:
		bank.saveAccount(ssn, new(accounts.CreditAccount).New(owner, interestRate, false))
		break
	case CreditCard:
		bank.saveAccount(ssn, new(accounts.CreditCardAccount).New(owner))
		break
	case Savings:
		bank.saveAccount(ssn, new(accounts.SavingsAccount).New(owner, interestRate))
		break
	}
}

func (bank Bank) createPackage(owner customer.Customer) {
	bank.createAccount(owner, Checking, 0)
	bank.createAccount(owner, Credit, 5)
	bank.createAccount(owner, CreditCard, 0)
	bank.createAccount(owner, Savings, 1)
}

func (bank Bank) saveAccount(ssn string, account accounts.IAccount) {
	bank.Accounts[ssn] = append(bank.Accounts[ssn], account)
}

func (bank Bank) computeInterest() {
	for _, acts := range bank.Accounts {
		for _, account := range acts {
			account.ComputeInterest()
		}
	}
}

func (bank Bank) findAccounts(ssn string, category string) []accounts.IAccount {
	for customerSsn, customerAccounts := range bank.Accounts {
		if customerSsn == ssn {
			if len(category) == 0 {
				return customerAccounts
			}

			var foundAccounts []accounts.IAccount
			for _, account := range customerAccounts {
				if account.Category() == category {
					foundAccounts = append(foundAccounts, account)
				}
			}

			return foundAccounts
		}
	}

	return nil
}

func (bank Bank) accountsWarning() []string {
	var warnings []string

	for _, customerAccounts := range bank.Accounts {
		for _, account := range customerAccounts {
			if account.Category() == Checking && account.Amount() < 0 {
				warnings = append(warnings, "Warning: "+account.Owner().Name+" has a negative balance for their checking account.")
			} else if account.Category() == CreditCard && account.Amount() < -5000 {
				warnings = append(warnings, "Warning: "+account.Owner().Name+" has an invalid balance for their credit card account.")
			}
		}
	}

	return warnings
}

func (bank Bank) totalDeposits() float32 {
	var total float32

	for _, customerAccounts := range bank.Accounts {
		for _, account := range customerAccounts {
			if account.Category() == Checking || account.Category() == Savings {
				if account.Amount() > 0 {
					total += account.Amount()
				}
			}
		}
	}

	return total
}

func (bank Bank) totalCredits() float32 {
	var total float32

	for _, customerAccounts := range bank.Accounts {
		for _, account := range customerAccounts {
			if account.Category() == Credit || account.Category() == CreditCard {
				total += account.Amount()
			}
		}
	}

	return total
}
