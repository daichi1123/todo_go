package test

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mike")
	want := "Hello, Mike"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
