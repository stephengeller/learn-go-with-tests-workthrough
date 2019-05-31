package internal

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assessResult := func(t *testing.T, got, want int, numbers []int) {
		if got != want {
			t.Errorf("got '%d', want '%d', given %v", got, want, numbers)
		}
	}

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v', want '%v'", got, want)
		}
	}

	t.Run("sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assessResult(t, got, want, numbers)
	})

	t.Run("any size", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6

		assessResult(t, got, want, nums)
	})

	t.Run("SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)

	})

	t.Run("SumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("SumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

	t.Run("somethi", func(t *testing.T) {
		got := MoreSliceStuff()
		want := []int{0, 0, 0}

		checkSums(t, got, want)
	})
}
