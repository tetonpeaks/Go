package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	assertCorrectMessage(t, repeated, expected)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	fmt.Printf("got: %v!\n", got)
}
