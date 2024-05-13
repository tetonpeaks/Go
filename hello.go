package main

import "fmt"

// main() is a package
/* func main() {
	fmt.Println("Hello, world")
} */

// It is good to separate your "domain" code from the outside world (side-effects).
// The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
// Func defined with string keyword, meaning a string is returned.
// Normally, as part of the TDD cycle, we should now refactor.
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
