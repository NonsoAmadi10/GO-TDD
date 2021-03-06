package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String()
}
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

func(w *Wallet) Deposit(amount Bitcoin)  {

	w.balance += amount

}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
var ErrMsg = errors.New("insufficient funds")
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrMsg
	}
	w.balance -= amount 
	return nil
}