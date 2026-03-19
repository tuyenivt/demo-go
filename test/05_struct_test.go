package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := rectangle.Perimeter()
	expected := 40.0

	if actual != expected {
		t.Errorf("expected %.2f but got %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		actual := rectangle.Area()
		expected := 72.0

		if actual != expected {
			t.Errorf("expected %.2f but got %.2f", expected, actual)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		actual := circle.Area()
		expected := 314.1592653589793

		if actual != expected {
			t.Errorf("expected %.2f but got %.2f", expected, actual)
		}
	})
}
