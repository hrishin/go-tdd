package main

import "testing"

func TestHelloWorld(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("test deposite", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposite(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalacne := Bitcoin(20)
		wallet := Wallet{startingBalacne}
		err := wallet.Withdraw(Bitcoin(30))

		assertBalance(t, wallet, startingBalacne)
		if err == nil {
			t.Error("wanted an error but didin't get")
		}
	})
}
