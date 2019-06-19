package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(bitcoin Bitcoin) {
	w.balance += bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {
	if bitcoin > w.Balance() {
		return errors.New("insufficient funds")
	}

	w.balance -= bitcoin
	return nil
}