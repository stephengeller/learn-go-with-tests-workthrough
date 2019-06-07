package internal

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		t.Helper()
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("didn't want an error but got one")
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit 10", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Deposit 20", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw 10", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)

		assertBalance(t, wallet, Bitcoin(90))
	})

	t.Run("Withdraw 20", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(20))

		assertBalance(t, wallet, Bitcoin(0))
		assertError(t, err, ErrInsufficientFunds)
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		startingBalance := Bitcoin(99)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(99))
		assertError(t, err, ErrInsufficientFunds)
	})
}
