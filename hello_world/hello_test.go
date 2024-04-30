package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Hello("John", "English")
		want := "Hello, John"

		assertCorrectMassege(t, got, want)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		got := Hello("John", "French")
		want := "Bonjour, John"

		assertCorrectMassege(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("John", "Spanish")
		want := "Hola, John"

		assertCorrectMassege(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		assertCorrectMassege(t, got, want)
	})
}

func assertCorrectMassege(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
