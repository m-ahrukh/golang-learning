package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	assert.Equal(t, want, got)
}

func TestWallets(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		assert.Equal(t, want, got)
		if got != want {
			t.Errorf("got %s \nwant %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// got := wallet.Balance()
		// want := Bitcoin(10)

		// assert.Equal(t, want, got)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		// got := wallet.Balance()

		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(30))
		err := wallet.Withdraw(Bitcoin(100))
		// got := wallet.Balance()

		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }

		assertBalance(t, wallet, Bitcoin(10))
		if err == nil {
			t.Error("wanted an error but didnt get one")
		}
	})
}
