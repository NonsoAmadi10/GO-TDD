package main

import (
	"testing"
)

func TestWallet(t *testing.T){

	t.Run("Deposit into wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got:= wallet.Balance()
		want := Bitcoin(10)

	


		if got != want {
			t.Errorf("got %s but wanted %s", got, want)
		}
	})

	t.Run("Withdraw from wallet", func(t *testing.T) {
			wallet := Wallet {}
			wallet.Deposit(Bitcoin(200))
			wallet.Withdraw(Bitcoin(50))

			got := wallet.Balance()
			want := Bitcoin(150)

			if got != want {
				t.Errorf("got %s but wanted %s", got, want)
			}
	})
}