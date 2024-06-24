package main

import (
	"errors"
	"fmt"
)

// Account struct to hold account information
type Account struct {
	AccountNumber string
	Balance       float64
}

// Bank struct to hold multiple accounts
type Bank struct {
	Accounts map[string]*Account
}

// NewBank creates a new Bank instance
func NewBank() *Bank {
	return &Bank{Accounts: make(map[string]*Account)}
}

// CreateAccount creates a new account with a given account number
func (b *Bank) CreateAccount(accountNumber string) (*Account, error) {
	if _, exists := b.Accounts[accountNumber]; exists {
		return nil, errors.New("account already exists")
	}
	account := &Account{AccountNumber: accountNumber, Balance: 0.0}
	b.Accounts[accountNumber] = account
	return account, nil
}

// Deposit adds money to the specified account
func (b *Bank) Deposit(accountNumber string, amount float64) error {
	account, exists := b.Accounts[accountNumber]
	if !exists {
		return errors.New("account not found")
	}
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	account.Balance += amount
	return nil
}

// Withdraw removes money from the specified account
func (b *Bank) Withdraw(accountNumber string, amount float64) error {
	account, exists := b.Accounts[accountNumber]
	if !exists {
		return errors.New("account not found")
	}
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > account.Balance {
		return errors.New("insufficient funds")
	}
	account.Balance -= amount
	return nil
}

// GetBalance returns the balance of the specified account
func (b *Bank) GetBalance(accountNumber string) (float64, error) {
	account, exists := b.Accounts[accountNumber]
	if !exists {
		return 0.0, errors.New("account not found")
	}
	return account.Balance, nil
}

func main() {
	bank := NewBank()
	bank.CreateAccount("123456")
	bank.Deposit("123456", 500.0)
	bank.Withdraw("123456", 200.0)
	balance, _ := bank.GetBalance("123456")
	fmt.Printf("Account 123456 balance: %.2f\n", balance)
}
