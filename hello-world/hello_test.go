package main

import "testing"

func TestHello(t *testing.T) {
	assertConnrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"
		assertConnrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertConnrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertConnrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Michael", "French")
		want := "Bonjour, Michael"
		assertConnrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Paul", "German")
		want := "Hallo, Paul"
		assertConnrectMessage(t, got, want)
	})
}
