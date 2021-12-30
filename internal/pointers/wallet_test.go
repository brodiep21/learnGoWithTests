package pointers

import (
	"reflect"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(140)}
		wallet.Withdraw(Bitcoin(30))
		assertBalance(t, wallet, Bitcoin(110))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(50)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
