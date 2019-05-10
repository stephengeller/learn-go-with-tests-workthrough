package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Stephen")
	want := "Hello, Stephen"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestWorksWithOtherNames(t *testing.T) {
	got := Hello("Dylan")
	want := "Hello, Dylan"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
