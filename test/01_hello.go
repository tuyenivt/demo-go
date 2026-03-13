package main

import "fmt"

const englishHelloFormatString = "Hello, %s"
const spanishHelloFormatString = "Hola, %s"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return fmt.Sprintf(spanishHelloFormatString, name)
	}
	return fmt.Sprintf(englishHelloFormatString, name)
}
