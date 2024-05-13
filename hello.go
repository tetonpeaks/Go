package main

import (
	"fmt"
)

// main() is a package
/* func main() {
	fmt.Println("Hello, world")
} */

// It is good to separate your "domain" code from the outside world (side-effects).
// The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
// Func defined with string keyword, meaning a string is returned.
// Normally, as part of the TDD cycle, we should now refactor.
// Refactored Hello

// We can group constants in a block instead of declaring them on their own line.
// For readability, it's a good idea to use a line between sets of related constants.
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

// (prefix string) is a named return value assigned the "zero" value [int 0 or string ""]
// In Go, public functions start with a capital letter, and private ones start with a lowercase letter.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
