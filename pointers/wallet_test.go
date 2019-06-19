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

		err := wallet.Withdraw(Bitcoin(5))

		assertNotError(t, err)
		assertBalance(t, wallet, Bitcoin(15))
	})
	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		// https://github.com/quii/learn-go-with-tests/blob/master/pointers-and-errors.md#write-the-test-first-2
		startBal := Bitcoin(20)
		wallet := Wallet{startBal}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startBal)
		assertError(t, err)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()

	actual := wallet.Balance()

	if actual != expected {
		t.Errorf("got %s want %s", actual, expected)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected an error bug didn't get one")
	}
}

func assertNotError(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Errorf("Unexpected Error: %s", e)
	}
}