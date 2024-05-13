// Go is cool!

package integers

import (
	"fmt"
	"strconv"
	"testing"
)

func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T) {
	t.Run("2+2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, strconv.Itoa(sum), strconv.Itoa(expected))
	})

	t.Run("1+5", func(t *testing.T) {
		sum := Add(1, 5)
		expected := 1 + 5
		assertCorrectMessage(t, strconv.Itoa(sum), strconv.Itoa(expected))
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	fmt.Printf("got: %v\n", got)
}
