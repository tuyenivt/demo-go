package main

import "fmt"

const helloFormatString = "Hello, %s"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return fmt.Sprintf(helloFormatString, name)
}
