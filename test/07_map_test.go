package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	actual := dictionary.Search("test")
	expected := "this is just a test"

	if actual != expected {
		t.Errorf("expected %q but got %q given %q", expected, actual, "test")
	}
}
