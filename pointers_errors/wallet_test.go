package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but did not get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("wanted an error but did not get one")
		}
	}
	t.Run("Testing deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})
	t.Run("Testing withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.balance = Bitcoin(20)
		err := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, got, want)
	})
	t.Run("Testing withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.balance = Bitcoin(20)
		err := wallet.Withdraw(Bitcoin(100))
		got := wallet.Balance()
		want := Bitcoin(20)
		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, got, want)
	})
}
