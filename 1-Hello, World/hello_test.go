package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Paulo")
	want := "Hello, Paulo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}