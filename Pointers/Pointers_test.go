package Pointers

import "testing"

func TestWalletDeposit(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, btc Bitcoin) {
		t.Helper()
		if wallet.balance != btc {
			t.Errorf("got %s, but want %s", wallet.balance, btc)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw no error", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(20))
		if err != nil {
			t.Fatal("Got an error when it was not expected")
		}
		assertBalance(t, wallet, Bitcoin(80))
	})

	t.Run("Withdraw with error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		err := wallet.Withdraw(Bitcoin(10))
		if err == nil {
			t.Fatal("Wanted an error but did not receive one")
		}
		if err.Error() != "Withdraw unsuccessful, insufficent funds" {
			t.Errorf("We got the wrong error message we wanted Withdraw unsucessful, insufficent funds but got %s", err.Error())
		}
		assertBalance(t, wallet, Bitcoin(5))
	})

}
