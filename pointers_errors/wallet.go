package pointerserrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficientFunds
	}
	wallet.balance -= amount
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
