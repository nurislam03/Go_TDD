package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of bitcoin
type Wallet struct {
	balance Bitcoin
}

// Deposit will add some amount of Bitcoin to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw subtracts some Bitcoin from the wallet, returning an error if it cann't be perform
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the number of Bitcoin a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {}