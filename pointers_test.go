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

	assertError := func (t testing.TB, got error, want error)  {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
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

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		
		assertError(t, err, ErrMsg)
		assertBalance(t, wallet, startingBalance)
	
		
	})
}