package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Max")
		expect := "Hello, Max"

		if actual != expect {
			t.Errorf("actual %q expect %q", actual, expect)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("")
		expect := "Hello, World"

		if actual != expect {
			t.Errorf("actual %q expect %q", actual, expect)
		}
	})
}
