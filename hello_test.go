package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	// f stands for format
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

//For now, it's enough to know that your t of type *testing.T is your "hook"
//into the testing framework so you can do things like t.Fail() when you want to fail.
