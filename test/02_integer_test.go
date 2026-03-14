package main

import "testing"

func TestAdder(t *testing.T) {
	actual := Adder(1, 2)
	expected := 3

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}
