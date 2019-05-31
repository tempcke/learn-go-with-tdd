package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(15))
	})
	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		t.Error("https://github.com/quii/learn-go-with-tests/blob/master/pointers-and-errors.md#write-the-test-first-2")
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()

	actual := wallet.Ballance()

	if actual != expected {
		t.Errorf("got %s want %s", actual, expected)
	}
}