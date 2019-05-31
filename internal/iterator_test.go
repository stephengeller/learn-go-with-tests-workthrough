package internal

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertGotExpected := func(t *testing.T, got, expected string) {
		if expected != got {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}

	t.Run("With any letter", func(t *testing.T) {
		expected := Repeat("a", 4)
		got := "aaaa"

		assertGotExpected(t, got, expected)
	})

	t.Run("With another letter", func(t *testing.T) {
		expected := Repeat("b", 4)
		got := "bbbb"

		assertGotExpected(t, got, expected)
	})

	t.Run("With another length", func(t *testing.T) {
		expected := Repeat("b", 8)
		got := "bbbbbbbb"

		assertGotExpected(t, got, expected)
	})
}

func TestExtraFns(t *testing.T) {

	assertGotExpected := func(t *testing.T, got, expected string) {
		if expected != got {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}

	t.Run("With some trim", func(t *testing.T) {
		expected := TrimSuffixOfString("some random string", "ing")
		got := "some random str"

		assertGotExpected(t, got, expected)
	})

	t.Run("With another trim", func(t *testing.T) {
		expected := "some r"
		got := TrimSuffixOfString("some random string", "andom string")

		assertGotExpected(t, got, expected)
	})

	t.Run("other fn", func(t *testing.T) {
		expected := "fooandbar"
		got := workOutTypesWithOnlyStrings("foo", "andbar")

		assertGotExpected(t, got, expected)
	})

	t.Run("other fn", func(t *testing.T) {
		expected := "45 is a number, blah is a string"
		got := workOutTypesWithMix(45, "blah")

		assertGotExpected(t, got, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 6)
	fmt.Printf(result)
	// Output: aaaaaa
}

func BenchmarkTrimSuffixOfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimSuffixOfString("bartholemew", "mew")
	}
}
