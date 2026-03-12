package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Max")
		expect := "Hello, Max"

		assertCorrectMessage(t, actual, expect)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("")
		expect := "Hello, World"

		assertCorrectMessage(t, actual, expect)
	})
}

func assertCorrectMessage(t testing.TB, actual, expect string) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual %q expect %q", actual, expect)
	}
}
