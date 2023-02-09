package repository

import (
	"fmt"
	"log"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("GetHelloWorld", func(t *testing.T) {
		got, _ := GetHelloWorld("test")
		want := "test"

		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleGetHelloWorld() {
	got, err := GetHelloWorld("test")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(got)
	// Output: test
}
