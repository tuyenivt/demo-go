package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := Perimeter(rectangle)
	expected := 40.0

	if actual != expected {
		t.Errorf("expected %.2f but got %.2f", expected, actual)
	}
}

func TestPArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	actual := Area(rectangle)
	expected := 72.0

	if actual != expected {
		t.Errorf("expected %.2f but got %.2f", expected, actual)
	}
}
