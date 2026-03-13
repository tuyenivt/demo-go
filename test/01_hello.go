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
	var formatString string
	switch language {
	case spanish:
		formatString = spanishHelloFormatString
	case french:
		formatString = frenchHelloFormatString
	default:
		formatString = englishHelloFormatString
	}
	return fmt.Sprintf(formatString, name)
}
