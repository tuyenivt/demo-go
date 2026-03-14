package main

import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 5)
	expected := "aaaaa"

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
