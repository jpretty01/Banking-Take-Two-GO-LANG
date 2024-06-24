package main

import (
	"testing"
)

func TestBank(t *testing.T) {
	bank := NewBank()

	// Test creating an account
	account, err := bank.CreateAccount("123456")
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}
	if account.AccountNumber != "123456" {
		t.Fatalf("Account number mismatch: expected %v, got %v", "123456", account.AccountNumber)
	}

	// Test depositing money
	err = bank.Deposit("123456", 1000.0)
	if err != nil {
		t.Fatalf("Failed to deposit money: %v", err)
	}
	balance, err := bank.GetBalance("123456")
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}
	if balance != 1000.0 {
		t.Fatalf("Balance mismatch: expected %v, got %v", 1000.0, balance)
	}

	// Test withdrawing money
	err = bank.Withdraw("123456", 500.0)
	if err != nil {
		t.Fatalf("Failed to withdraw money: %v", err)
	}
	balance, err = bank.GetBalance("123456")
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}
	if balance != 500.0 {
		t.Fatalf("Balance mismatch: expected %v, got %v", 500.0, balance)
	}

	// Test withdrawing more money than the balance
	err = bank.Withdraw("123456", 1000.0)
	if err == nil {
		t.Fatalf("Expected error when withdrawing more than the balance, got nil")
	}

	// Test depositing negative amount
	err = bank.Deposit("123456", -100.0)
	if err == nil {
		t.Fatalf("Expected error when depositing negative amount, got nil")
	}

	// Test withdrawing negative amount
	err = bank.Withdraw("123456", -100.0)
	if err == nil {
		t.Fatalf("Expected error when withdrawing negative amount, got nil")
	}

	// Test creating duplicate account
	_, err = bank.CreateAccount("123456")
	if err == nil {
		t.Fatalf("Expected error when creating duplicate account, got nil")
	}

	// Test getting balance of non-existent account
	_, err = bank.GetBalance("000000")
	if err == nil {
		t.Fatalf("Expected error when getting balance of non-existent account, got nil")
	}
}
