package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Should hello to people", func(t *testing.T) {
		got := Hello("Paul")
		want := "Hello, Paulo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should use World when input is empty", func (t *testing.T)  {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

