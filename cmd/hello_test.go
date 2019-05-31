package cmd

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Logf("got '%s', want '%s'", got, want)
			t.Fail()
		}
	}

	t.Run("WithMyName", func(t *testing.T) {
		got := Hello("Stephen", "English")
		want := "Hello, Stephen"
		assertCorrectMessage(t, got, want)
	})

	t.Run("WorksWithOtherNames", func(t *testing.T) {
		got := Hello("Dylan", "English")
		want := "Hello, Dylan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("WorksWithNoName", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Stephen", "Spanish")
		want := "Hola, Stephen"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Stephen", "French")
		want := "Bonjour, Stephen"
		assertCorrectMessage(t, got, want)
	})

}
