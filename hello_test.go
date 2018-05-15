package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func TestHelloWithArgument(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Saying hello to folks", func(t *testing.T) {
		got := HelloString("Hrishi")
		want := "Hello, Hrishi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello World when empty string is suplied", func(t *testing.T) {
		got := HelloString("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}
