package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Maria", "Spanish")
		want := "Hola, Maria"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Madamoiselle", "French")
		want := "Bonjour, Madamoiselle"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	fmt.Printf("got: %v\n", got)
}

//t.Helper() is needed to tell the test suite that this method is a helper. By doing this, when it fails, the line number
//reported will be in our function call rather than inside our test helper. This will help other developers track down problems more easily.

//For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B
//both satisfy, so you can call helper functions from a test, or a benchmark

//We've refactored our assertion into a new function. This reduces duplication and
//improves the readability of our tests. We need to pass in t *testing.T so that we
//can tell the test code to fail when we need to.

//Here, we are introducing another tool in our testing arsenal: subtests.
//Sometimes, it is useful to group tests around a "thing" and then have subtests describing different scenarios.
//A benefit of this approach is you can set up shared code that can be used in the other tests.
//It is important that your tests are clear specifications of what the code needs to do. But there is repeated code
//when we check if the message is what we expect.

//For now, it's enough to know that your t of type *testing.T is your "hook"
//into the testing framework so you can do things like t.Fail() when you want to fail.
