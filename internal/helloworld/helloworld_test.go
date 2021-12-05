package helloworld_test

import (
	"testing"

	"github.com/brodiep21/learnGoWithTests/internal/helloworld"
)

func TestHello(t *testing.T) {
	t.Run("HelloBrodie", func(t *testing.T) {
		got := helloworld.Hello("Brodie")
		want := "Hello, Brodie"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("HelloEli", func(t *testing.T) {
		got := helloworld.Hello("Eli")
		want := "Hello, Eli"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("HelloNobody", func(t *testing.T) {
		got := helloworld.Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
