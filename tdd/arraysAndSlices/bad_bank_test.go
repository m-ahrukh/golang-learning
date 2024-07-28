package arraysandslices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	got := int(newBalanceFor(riya))
	assert.Equal(t, got, 200)
	got = int(newBalanceFor(chris))
	assert.Equal(t, got, 0)
	got = int(newBalanceFor(adil))
	assert.Equal(t, got, 175)
}
