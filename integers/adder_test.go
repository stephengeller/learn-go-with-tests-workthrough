package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCorrectSum := func(t *testing.T, sum, expected int) {
		if sum != expected {
			t.Errorf("Expected '%d', but got '%d'", expected, sum)
		}
	}
	t.Run("with number", func(t *testing.T) {
		sum := Add(1, 7)
		expected := 8

		if sum != expected {
			t.Errorf("Expected '%d', but got '%d'", expected, sum)
		}
	})

	t.Run("with any numbers", func(t *testing.T) {
		sum := Add(4, 1)
		expected := 5
		assertCorrectSum(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(5, 5)
	fmt.Println(sum)
	// Output: 10

}
