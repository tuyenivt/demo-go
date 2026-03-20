package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		actual := wallet.Balance()
		expected := startingBalance

		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}

		if err == nil {
			t.Error("expected error but didn't get one")
		}
	})
}
