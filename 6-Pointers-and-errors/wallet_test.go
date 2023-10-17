package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		expected := Bitcoin(10)

		if got != expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	})
}