package main

import "fmt"

const helloFormatString = "Hello, %s"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(helloFormatString, name)
}
