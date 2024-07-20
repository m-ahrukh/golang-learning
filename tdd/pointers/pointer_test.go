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

	// assertError := func(t testing.TB, err error) {
	// 	t.Helper()
	// 	if err == nil {
	// 		t.Error("Wanted an error but didn't get one")
	// 	}
	// }

	// assertError := func(t testing.TB, got error, want string) {
	// 	t.Helper()
	// 	if got == nil {
	// 		t.Fatal("didn't get an error but wanted one")
	// 	}
	// 	if got.Error() != want {
	// 		t.Errorf("\ngot %q, \nwant %q", got, want)
	// 	}
	// }

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("\ngot %q, \nwant %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
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

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		// got := wallet.Balance()
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(30))
		// got := wallet.Balance()
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

	t.Run("withdraw insufficient funds 2.0", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))
		// got := wallet.Balance()
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}
