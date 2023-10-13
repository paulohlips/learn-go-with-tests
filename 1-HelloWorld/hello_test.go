package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Should hello to people", func(t *testing.T) {
		got := Hello("User", "")
		want := "Hello, User"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should use World when input is empty", func (t *testing.T)  {
		got := Hello("","")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should use spanish", func(t *testing.T) {
		got := Hello("User", "Spanish")
		want := "Hola, User"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should use french", func(t *testing.T) {
		got := Hello("User", "French")
		want := "Bonjour, User"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should use portuguese", func(t *testing.T) {
		got := Hello("User", "Portuguese")
		want := "Olá, User"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

