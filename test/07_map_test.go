package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		if actual != expected {
			t.Errorf("expected %q but got %q given %q", expected, actual, "test")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		if err.Error() != expected {
			t.Errorf("expected %q but got %q given %q", expected, err.Error(), "unknown")
		}
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	actual, err := dictionary.Search("test")
	expected := "this is just a test"

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if actual != expected {
		t.Errorf("expected %q but got %q given %q", expected, actual, "test")
	}
}
