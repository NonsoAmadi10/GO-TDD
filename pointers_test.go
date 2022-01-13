package main

import (
	"testing"
)

func TestWallet(t *testing.T){

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    }

	t.Run("Deposit into wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw from wallet", func(t *testing.T) {
			wallet := Wallet {}
			wallet.Deposit(Bitcoin(200))
			wallet.Withdraw(Bitcoin(50))
			want := Bitcoin(150)
			assertBalance(t, wallet, want)
		
	})
}