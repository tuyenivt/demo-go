package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("expected %d but got %d, %v", expected, actual, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("expected %d but got %d, %v", expected, actual, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
