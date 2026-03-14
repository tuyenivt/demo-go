package main

import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat("a")
	expected := "aaaaa"

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
