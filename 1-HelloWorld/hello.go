package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("User", ""))
}

const (
	french = "French"
	spanish = "Spanish"

	portuguese = "Portuguese"
	englishHelloPrefix = "Hello, "
	spanishHolaPrefix = "Hola, "
	frenchBounjourPrefix = "Bonjour, "
	portugueseOlaPrefix = "Olá, "
)

func Hello(name string, language string) string{
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchBounjourPrefix
	case spanish:
		prefix = spanishHolaPrefix
	case portuguese:
		prefix = portugueseOlaPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}