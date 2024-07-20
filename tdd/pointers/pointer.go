package pointers

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// func (w *Wallet) Withdraw(amount Bitcoin) {
// 	w.balance -= amount
// }

// func (w *Wallet) Withdraw(amount Bitcoin) error {
// 	w.balance -= amount
// 	return nil
// }

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// return errors.New("cannot withdraw, insufficient funds")
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
