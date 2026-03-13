package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloFormatString = "Hello, %s"
const spanishHelloFormatString = "Hola, %s"
const frenchHelloFormatString = "Bonjour, %s"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return fmt.Sprintf(spanishHelloFormatString, name)
	}
	if language == french {
		return fmt.Sprintf(frenchHelloFormatString, name)
	}
	return fmt.Sprintf(englishHelloFormatString, name)
}
