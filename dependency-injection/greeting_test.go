package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	// bytes.Buffer as method Write(b []byte) n int, err Error
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Alice")

	got := buffer.String()
	want := "Hello, Alice"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
