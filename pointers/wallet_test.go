package main

import "testing"

func TestHelloWorld(t *testing.T) {

	t.Run("test deposite", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposite(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
		assertNoErr(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalacne := Bitcoin(20)
		wallet := Wallet{startingBalacne}
		err := wallet.Withdraw(Bitcoin(30))

		assertBalance(t, wallet, startingBalacne)
		assertError(t, err, insufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Error("didin't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Error("got an error but didnt want one")
	}
}
