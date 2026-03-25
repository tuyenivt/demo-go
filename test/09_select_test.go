package main

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "http://www.facebook.com"
	fastUrl := "http://www.quii.dev"

	expected := fastUrl
	actual := Racer(slowUrl, fastUrl)

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
