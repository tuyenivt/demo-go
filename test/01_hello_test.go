package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("world")
	expect := "Hello, world"

	if actual != expect {
		t.Errorf("actual %q expect %q", actual, expect)
	}
}
