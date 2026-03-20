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
}
