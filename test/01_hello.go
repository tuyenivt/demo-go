package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloFormatString = "Hello, %s"
	spanishHelloFormatString = "Hola, %s"
	frenchHelloFormatString  = "Bonjour, %s"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(getFormatString(language), name)
}

func getFormatString(language string) string {
	switch language {
	case spanish:
		return spanishHelloFormatString
	case french:
		return frenchHelloFormatString
	default:
		return englishHelloFormatString
	}
}
