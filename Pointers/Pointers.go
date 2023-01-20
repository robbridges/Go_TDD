package Pointers

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin float64

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt

}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	var InsufficentFundsError = errors.New("Withdraw unsuccessful, insufficent funds")
	if w.balance < amt {
		return InsufficentFundsError
	}
	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// decent method to return our bitcoin to string, so that it's value can be passed as a string as well
func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
