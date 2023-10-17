package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		got := wallet.Balance()

		if got != expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		wallet.Balance()
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		wallet.Balance()
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)

	})
}